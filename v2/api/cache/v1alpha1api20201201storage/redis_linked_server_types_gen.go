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

// +kubebuilder:rbac:groups=cache.azure.com,resources=redislinkedservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.azure.com,resources={redislinkedservers/status,redislinkedservers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201201.RedisLinkedServer
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_linkedServers
type RedisLinkedServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisLinkedServers_Spec                `json:"spec,omitempty"`
	Status            RedisLinkedServerWithProperties_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisLinkedServer{}

// GetConditions returns the conditions of the resource
func (redisLinkedServer *RedisLinkedServer) GetConditions() conditions.Conditions {
	return redisLinkedServer.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (redisLinkedServer *RedisLinkedServer) SetConditions(conditions conditions.Conditions) {
	redisLinkedServer.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &RedisLinkedServer{}

// AzureName returns the Azure name of the resource
func (redisLinkedServer *RedisLinkedServer) AzureName() string {
	return redisLinkedServer.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (redisLinkedServer RedisLinkedServer) GetAPIVersion() string {
	return "2020-12-01"
}

// GetResourceKind returns the kind of the resource
func (redisLinkedServer *RedisLinkedServer) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (redisLinkedServer *RedisLinkedServer) GetSpec() genruntime.ConvertibleSpec {
	return &redisLinkedServer.Spec
}

// GetStatus returns the status of this resource
func (redisLinkedServer *RedisLinkedServer) GetStatus() genruntime.ConvertibleStatus {
	return &redisLinkedServer.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (redisLinkedServer *RedisLinkedServer) GetType() string {
	return "Microsoft.Cache/redis/linkedServers"
}

// NewEmptyStatus returns a new empty (blank) status
func (redisLinkedServer *RedisLinkedServer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisLinkedServerWithProperties_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (redisLinkedServer *RedisLinkedServer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(redisLinkedServer.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  redisLinkedServer.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (redisLinkedServer *RedisLinkedServer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisLinkedServerWithProperties_Status); ok {
		redisLinkedServer.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisLinkedServerWithProperties_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	redisLinkedServer.Status = st
	return nil
}

// Hub marks that this RedisLinkedServer is the hub type for conversion
func (redisLinkedServer *RedisLinkedServer) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (redisLinkedServer *RedisLinkedServer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: redisLinkedServer.Spec.OriginalVersion,
		Kind:    "RedisLinkedServer",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201201.RedisLinkedServer
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_linkedServers
type RedisLinkedServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisLinkedServer `json:"items"`
}

//Storage version of v1alpha1api20201201.RedisLinkedServerWithProperties_Status
type RedisLinkedServerWithProperties_Status struct {
	Conditions               []conditions.Condition `json:"conditions,omitempty"`
	Id                       *string                `json:"id,omitempty"`
	LinkedRedisCacheId       *string                `json:"linkedRedisCacheId,omitempty"`
	LinkedRedisCacheLocation *string                `json:"linkedRedisCacheLocation,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState        *string                `json:"provisioningState,omitempty"`
	ServerRole               *string                `json:"serverRole,omitempty"`
	Type                     *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisLinkedServerWithProperties_Status{}

// ConvertStatusFrom populates our RedisLinkedServerWithProperties_Status from the provided source
func (redisLinkedServerWithPropertiesStatus *RedisLinkedServerWithProperties_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == redisLinkedServerWithPropertiesStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(redisLinkedServerWithPropertiesStatus)
}

// ConvertStatusTo populates the provided destination from our RedisLinkedServerWithProperties_Status
func (redisLinkedServerWithPropertiesStatus *RedisLinkedServerWithProperties_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == redisLinkedServerWithPropertiesStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(redisLinkedServerWithPropertiesStatus)
}

//Storage version of v1alpha1api20201201.RedisLinkedServers_Spec
type RedisLinkedServers_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                string  `json:"azureName"`
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// +kubebuilder:validation:Required
	//LinkedRedisCacheReference: Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheReference genruntime.ResourceReference `armReference:"LinkedRedisCacheId" json:"linkedRedisCacheReference"`
	Location                  *string                      `json:"location,omitempty"`
	OriginalVersion           string                       `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner" kind:"Redis"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	ServerRole  *string                           `json:"serverRole,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisLinkedServers_Spec{}

// ConvertSpecFrom populates our RedisLinkedServers_Spec from the provided source
func (redisLinkedServersSpec *RedisLinkedServers_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == redisLinkedServersSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(redisLinkedServersSpec)
}

// ConvertSpecTo populates the provided destination from our RedisLinkedServers_Spec
func (redisLinkedServersSpec *RedisLinkedServers_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == redisLinkedServersSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(redisLinkedServersSpec)
}

func init() {
	SchemeBuilder.Register(&RedisLinkedServer{}, &RedisLinkedServerList{})
}
