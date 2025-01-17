/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Storage_StorageAccount_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Custom namer because storage accounts have strict names
	namer := tc.Namer.WithSeparator("")

	// Create a storage account
	accessTier := storage.StorageAccountPropertiesCreateParametersAccessTierHot
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(namer.GenerateName("stor")),
		Spec: storage.StorageAccounts_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     storage.StorageAccountsSpecKindBlobStorage,
			Sku: storage.Sku{
				Name: storage.SkuNameStandardLRS,
			},
			// TODO: They mark this property as optional but actually it is required
			AccessTier: &accessTier,
		},
	}

	tc.CreateResourceAndWait(acct)

	tc.Expect(acct.Status.Location).To(Equal(&tc.AzureRegion))
	expectedKind := storage.StorageAccountStatusKindBlobStorage
	tc.Expect(acct.Status.Kind).To(Equal(&expectedKind))
	tc.Expect(acct.Status.Id).ToNot(BeNil())
	armId := *acct.Status.Id

	// Run sub-tests on storage account
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Blob Services CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				StorageAccount_BlobServices_CRUD(tc, acct)
			},
		},
	)

	tc.DeleteResourceAndWait(acct)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(storage.StorageAccountsSpecAPIVersion20210401))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func StorageAccount_BlobServices_CRUD(tc *testcommon.KubePerTestContext, storageAccount client.Object) {
	blobService := &storage.StorageAccountsBlobService{
		ObjectMeta: tc.MakeObjectMeta("blobservice"),
		Spec: storage.StorageAccountsBlobServices_Spec{
			Owner: testcommon.AsOwner(storageAccount),
		},
	}

	tc.CreateResourceAndWait(blobService)
	// no DELETE, this is not a real resource

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Container CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				StorageAccount_BlobServices_Container_CRUD(testContext, blobService)
			},
		})

	// TODO: Delete doesn't seem to work?
	// — is this because it is not a real resource but properties on the storage account?
	// We probably need to ask the Storage team.
	/*
		err = testContext.KubeClient.Delete(ctx, blobService)
		g.Expect(err).ToNot(HaveOccurred())
		g.Eventually(blobService).Should(testContext.Match.BeDeleted())
	*/
}

func StorageAccount_BlobServices_Container_CRUD(tc *testcommon.KubePerTestContext, blobService client.Object) {
	blobContainer := &storage.StorageAccountsBlobServicesContainer{
		ObjectMeta: tc.MakeObjectMeta("container"),
		Spec: storage.StorageAccountsBlobServicesContainers_Spec{
			Owner: testcommon.AsOwner(blobService),
		},
	}

	tc.CreateResourceAndWait(blobContainer)
	defer tc.DeleteResourceAndWait(blobContainer)
}
