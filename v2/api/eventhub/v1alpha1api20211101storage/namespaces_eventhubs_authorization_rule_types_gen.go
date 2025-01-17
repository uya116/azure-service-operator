// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventhub.azure.com,resources=namespaceseventhubsauthorizationrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventhub.azure.com,resources={namespaceseventhubsauthorizationrules/status,namespaceseventhubsauthorizationrules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20211101.NamespacesEventhubsAuthorizationRule
//Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs_authorizationRules
type NamespacesEventhubsAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesEventhubsAuthorizationRules_Spec `json:"spec,omitempty"`
	Status            AuthorizationRule_Status                   `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesEventhubsAuthorizationRule{}

// GetConditions returns the conditions of the resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) GetConditions() conditions.Conditions {
	return namespacesEventhubsAuthorizationRule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) SetConditions(conditions conditions.Conditions) {
	namespacesEventhubsAuthorizationRule.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NamespacesEventhubsAuthorizationRule{}

// AzureName returns the Azure name of the resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) AzureName() string {
	return namespacesEventhubsAuthorizationRule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (namespacesEventhubsAuthorizationRule NamespacesEventhubsAuthorizationRule) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceKind returns the kind of the resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) GetSpec() genruntime.ConvertibleSpec {
	return &namespacesEventhubsAuthorizationRule.Spec
}

// GetStatus returns the status of this resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) GetStatus() genruntime.ConvertibleStatus {
	return &namespacesEventhubsAuthorizationRule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &AuthorizationRule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(namespacesEventhubsAuthorizationRule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  namespacesEventhubsAuthorizationRule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*AuthorizationRule_Status); ok {
		namespacesEventhubsAuthorizationRule.Status = *st
		return nil
	}

	// Convert status to required version
	var st AuthorizationRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	namespacesEventhubsAuthorizationRule.Status = st
	return nil
}

// Hub marks that this NamespacesEventhubsAuthorizationRule is the hub type for conversion
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: namespacesEventhubsAuthorizationRule.Spec.OriginalVersion,
		Kind:    "NamespacesEventhubsAuthorizationRule",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20211101.NamespacesEventhubsAuthorizationRule
//Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs_authorizationRules
type NamespacesEventhubsAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsAuthorizationRule `json:"items"`
}

//Storage version of v1alpha1api20211101.NamespacesEventhubsAuthorizationRules_Spec
type NamespacesEventhubsAuthorizationRules_Spec struct {
	// +kubebuilder:validation:MinLength=1
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string  `json:"azureName"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner" kind:"NamespacesEventhub"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Rights      []string                          `json:"rights,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesEventhubsAuthorizationRules_Spec{}

// ConvertSpecFrom populates our NamespacesEventhubsAuthorizationRules_Spec from the provided source
func (namespacesEventhubsAuthorizationRulesSpec *NamespacesEventhubsAuthorizationRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == namespacesEventhubsAuthorizationRulesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(namespacesEventhubsAuthorizationRulesSpec)
}

// ConvertSpecTo populates the provided destination from our NamespacesEventhubsAuthorizationRules_Spec
func (namespacesEventhubsAuthorizationRulesSpec *NamespacesEventhubsAuthorizationRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == namespacesEventhubsAuthorizationRulesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(namespacesEventhubsAuthorizationRulesSpec)
}

func init() {
	SchemeBuilder.Register(&NamespacesEventhubsAuthorizationRule{}, &NamespacesEventhubsAuthorizationRuleList{})
}
