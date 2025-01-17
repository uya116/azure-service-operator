// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cache.azure.com,resources=redisfirewallrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.azure.com,resources={redisfirewallrules/status,redisfirewallrules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201201.RedisFirewallRule
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_firewallRules
type RedisFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisFirewallRules_Spec  `json:"spec,omitempty"`
	Status            RedisFirewallRule_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisFirewallRule{}

// GetConditions returns the conditions of the resource
func (redisFirewallRule *RedisFirewallRule) GetConditions() conditions.Conditions {
	return redisFirewallRule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (redisFirewallRule *RedisFirewallRule) SetConditions(conditions conditions.Conditions) {
	redisFirewallRule.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &RedisFirewallRule{}

// AzureName returns the Azure name of the resource
func (redisFirewallRule *RedisFirewallRule) AzureName() string {
	return redisFirewallRule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (redisFirewallRule RedisFirewallRule) GetAPIVersion() string {
	return "2020-12-01"
}

// GetResourceKind returns the kind of the resource
func (redisFirewallRule *RedisFirewallRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (redisFirewallRule *RedisFirewallRule) GetSpec() genruntime.ConvertibleSpec {
	return &redisFirewallRule.Spec
}

// GetStatus returns the status of this resource
func (redisFirewallRule *RedisFirewallRule) GetStatus() genruntime.ConvertibleStatus {
	return &redisFirewallRule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/firewallRules"
func (redisFirewallRule *RedisFirewallRule) GetType() string {
	return "Microsoft.Cache/redis/firewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (redisFirewallRule *RedisFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisFirewallRule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (redisFirewallRule *RedisFirewallRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(redisFirewallRule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  redisFirewallRule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (redisFirewallRule *RedisFirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisFirewallRule_Status); ok {
		redisFirewallRule.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisFirewallRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	redisFirewallRule.Status = st
	return nil
}

// Hub marks that this RedisFirewallRule is the hub type for conversion
func (redisFirewallRule *RedisFirewallRule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (redisFirewallRule *RedisFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: redisFirewallRule.Spec.OriginalVersion,
		Kind:    "RedisFirewallRule",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201201.RedisFirewallRule
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_firewallRules
type RedisFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisFirewallRule `json:"items"`
}

//Storage version of v1alpha1api20201201.RedisFirewallRule_Status
type RedisFirewallRule_Status struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	EndIP       *string                `json:"endIP,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartIP     *string                `json:"startIP,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisFirewallRule_Status{}

// ConvertStatusFrom populates our RedisFirewallRule_Status from the provided source
func (redisFirewallRuleStatus *RedisFirewallRule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == redisFirewallRuleStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(redisFirewallRuleStatus)
}

// ConvertStatusTo populates the provided destination from our RedisFirewallRule_Status
func (redisFirewallRuleStatus *RedisFirewallRule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == redisFirewallRuleStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(redisFirewallRuleStatus)
}

//Storage version of v1alpha1api20201201.RedisFirewallRules_Spec
type RedisFirewallRules_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string  `json:"azureName"`
	EndIP           *string `json:"endIP,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner" kind:"Redis"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	StartIP     *string                           `json:"startIP,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisFirewallRules_Spec{}

// ConvertSpecFrom populates our RedisFirewallRules_Spec from the provided source
func (redisFirewallRulesSpec *RedisFirewallRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == redisFirewallRulesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(redisFirewallRulesSpec)
}

// ConvertSpecTo populates the provided destination from our RedisFirewallRules_Spec
func (redisFirewallRulesSpec *RedisFirewallRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == redisFirewallRulesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(redisFirewallRulesSpec)
}

func init() {
	SchemeBuilder.Register(&RedisFirewallRule{}, &RedisFirewallRuleList{})
}
