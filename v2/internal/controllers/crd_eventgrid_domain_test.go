/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/go-autorest/autorest/to"

	eventgrid "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_EventGrid_Domain(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	publicNetworkAccess := eventgrid.DomainPropertiesPublicNetworkAccessEnabled

	// Create a domain
	domain := &eventgrid.Domain{
		ObjectMeta: tc.MakeObjectMeta("domain"),
		Spec: eventgrid.Domains_Spec{
			Location:            tc.AzureRegion,
			Owner:               testcommon.AsOwner(rg),
			PublicNetworkAccess: &publicNetworkAccess,
		},
	}

	// Create a storage account to use as destination
	accessTier := storage.StorageAccountPropertiesCreateParametersAccessTierHot
	acctName := "dest"
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMeta(acctName),
		Spec: storage.StorageAccounts_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     storage.StorageAccountsSpecKindStorageV2,
			Sku: storage.Sku{
				Name: storage.SkuNameStandardLRS,
			},
			// TODO: They mark this property as optional but actually it is required
			AccessTier: &accessTier,
		},
	}

	// TODO: Getting this is SUPER awkward
	accountARMID, err := genericarmclient.MakeResourceGroupScopeARMID(
		tc.AzureSubscription,
		rg.Name,
		"Microsoft.Storage",
		"storageAccounts",
		acctName)
	if err != nil {
		panic(err)
	}

	acctReference := genruntime.ResourceReference{ARMID: accountARMID}

	tc.CreateResourcesAndWait(domain, acct)

	queueServices := &storage.StorageAccountsQueueService{
		ObjectMeta: tc.MakeObjectMeta("dest-queues"),
		Spec: storage.StorageAccountsQueueServices_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}

	tc.CreateResourceAndWait(queueServices)

	queueName := "dest-queue"
	queue := &storage.StorageAccountsQueueServicesQueue{
		ObjectMeta: tc.MakeObjectMeta(queueName),
		Spec: storage.StorageAccountsQueueServicesQueues_Spec{
			Owner: testcommon.AsOwner(queueServices),
		},
	}

	tc.CreateResourceAndWait(queue)

	armId := *domain.Status.Id
	// objectKey := client.ObjectKeyFromObject(domain)

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CreateDomainTopicAndSubscription",
			Test: func(tc *testcommon.KubePerTestContext) {
				topic := &eventgrid.DomainsTopic{
					ObjectMeta: tc.MakeObjectMeta("topic"),
					Spec: eventgrid.DomainsTopics_Spec{
						Owner: testcommon.AsOwner(domain),
					},
				}

				tc.CreateResourceAndWait(topic)
				// don’t bother deleting; deleting domain will clean up

				subscription := &eventgrid.EventSubscription{
					ObjectMeta: tc.MakeObjectMeta("sub"),
					Spec: eventgrid.EventSubscriptions_Spec{
						Owner: tc.AsExtensionOwner(topic),
						Destination: &eventgrid.EventSubscriptionDestination{
							StorageQueue: &eventgrid.StorageQueueEventSubscriptionDestination{
								EndpointType: eventgrid.StorageQueueEventSubscriptionDestinationEndpointTypeStorageQueue,
								Properties: &eventgrid.StorageQueueEventSubscriptionDestinationProperties{
									ResourceReference: &acctReference,
									QueueName:         to.StringPtr(queueName),
								},
							},
						},
					},
				}

				tc.CreateResourceAndWait(subscription)
				// don’t bother deleting
			},
		},
		testcommon.Subtest{
			Name: "CreateDomainSubscription",
			Test: func(tc *testcommon.KubePerTestContext) {
				subscription := &eventgrid.EventSubscription{
					ObjectMeta: tc.MakeObjectMeta("sub"),
					Spec: eventgrid.EventSubscriptions_Spec{
						Owner: tc.AsExtensionOwner(domain),
						Destination: &eventgrid.EventSubscriptionDestination{
							StorageQueue: &eventgrid.StorageQueueEventSubscriptionDestination{
								EndpointType: eventgrid.StorageQueueEventSubscriptionDestinationEndpointTypeStorageQueue,
								Properties: &eventgrid.StorageQueueEventSubscriptionDestinationProperties{
									ResourceReference: &acctReference,
									QueueName:         to.StringPtr(queueName),
								},
							},
						},
					},
				}

				tc.CreateResourceAndWait(subscription)
				// don’t bother deleting
			},
		},
	)

	tc.DeleteResourceAndWait(domain)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(eventgrid.TopicsSpecAPIVersion20200601))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
