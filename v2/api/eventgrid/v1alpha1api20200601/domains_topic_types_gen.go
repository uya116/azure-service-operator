// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/resourceDefinitions/domains_topics
type DomainsTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainsTopics_Spec `json:"spec,omitempty"`
	Status            DomainTopic_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DomainsTopic{}

// GetConditions returns the conditions of the resource
func (domainsTopic *DomainsTopic) GetConditions() conditions.Conditions {
	return domainsTopic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (domainsTopic *DomainsTopic) SetConditions(conditions conditions.Conditions) {
	domainsTopic.Status.Conditions = conditions
}

var _ conversion.Convertible = &DomainsTopic{}

// ConvertFrom populates our DomainsTopic from the provided hub DomainsTopic
func (domainsTopic *DomainsTopic) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20200601storage.DomainsTopic)
	if !ok {
		return fmt.Errorf("expected storage:eventgrid/v1alpha1api20200601storage/DomainsTopic but received %T instead", hub)
	}

	return domainsTopic.AssignPropertiesFromDomainsTopic(source)
}

// ConvertTo populates the provided hub DomainsTopic from our DomainsTopic
func (domainsTopic *DomainsTopic) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20200601storage.DomainsTopic)
	if !ok {
		return fmt.Errorf("expected storage:eventgrid/v1alpha1api20200601storage/DomainsTopic but received %T instead", hub)
	}

	return domainsTopic.AssignPropertiesToDomainsTopic(destination)
}

// +kubebuilder:webhook:path=/mutate-eventgrid-azure-com-v1alpha1api20200601-domainstopic,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventgrid.azure.com,resources=domainstopics,verbs=create;update,versions=v1alpha1api20200601,name=default.v1alpha1api20200601.domainstopics.eventgrid.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &DomainsTopic{}

// Default applies defaults to the DomainsTopic resource
func (domainsTopic *DomainsTopic) Default() {
	domainsTopic.defaultImpl()
	var temp interface{} = domainsTopic
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (domainsTopic *DomainsTopic) defaultAzureName() {
	if domainsTopic.Spec.AzureName == "" {
		domainsTopic.Spec.AzureName = domainsTopic.Name
	}
}

// defaultImpl applies the code generated defaults to the DomainsTopic resource
func (domainsTopic *DomainsTopic) defaultImpl() { domainsTopic.defaultAzureName() }

var _ genruntime.KubernetesResource = &DomainsTopic{}

// AzureName returns the Azure name of the resource
func (domainsTopic *DomainsTopic) AzureName() string {
	return domainsTopic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (domainsTopic DomainsTopic) GetAPIVersion() string {
	return "2020-06-01"
}

// GetResourceKind returns the kind of the resource
func (domainsTopic *DomainsTopic) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (domainsTopic *DomainsTopic) GetSpec() genruntime.ConvertibleSpec {
	return &domainsTopic.Spec
}

// GetStatus returns the status of this resource
func (domainsTopic *DomainsTopic) GetStatus() genruntime.ConvertibleStatus {
	return &domainsTopic.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (domainsTopic *DomainsTopic) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (domainsTopic *DomainsTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DomainTopic_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (domainsTopic *DomainsTopic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(domainsTopic.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  domainsTopic.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (domainsTopic *DomainsTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DomainTopic_Status); ok {
		domainsTopic.Status = *st
		return nil
	}

	// Convert status to required version
	var st DomainTopic_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	domainsTopic.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-eventgrid-azure-com-v1alpha1api20200601-domainstopic,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventgrid.azure.com,resources=domainstopics,verbs=create;update,versions=v1alpha1api20200601,name=validate.v1alpha1api20200601.domainstopics.eventgrid.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &DomainsTopic{}

// ValidateCreate validates the creation of the resource
func (domainsTopic *DomainsTopic) ValidateCreate() error {
	validations := domainsTopic.createValidations()
	var temp interface{} = domainsTopic
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (domainsTopic *DomainsTopic) ValidateDelete() error {
	validations := domainsTopic.deleteValidations()
	var temp interface{} = domainsTopic
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (domainsTopic *DomainsTopic) ValidateUpdate(old runtime.Object) error {
	validations := domainsTopic.updateValidations()
	var temp interface{} = domainsTopic
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (domainsTopic *DomainsTopic) createValidations() []func() error {
	return []func() error{domainsTopic.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (domainsTopic *DomainsTopic) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (domainsTopic *DomainsTopic) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return domainsTopic.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (domainsTopic *DomainsTopic) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&domainsTopic.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromDomainsTopic populates our DomainsTopic from the provided source DomainsTopic
func (domainsTopic *DomainsTopic) AssignPropertiesFromDomainsTopic(source *v1alpha1api20200601storage.DomainsTopic) error {

	// ObjectMeta
	domainsTopic.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DomainsTopics_Spec
	err := spec.AssignPropertiesFromDomainsTopicsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesFromDomainsTopicsSpec()")
	}
	domainsTopic.Spec = spec

	// Status
	var status DomainTopic_Status
	err = status.AssignPropertiesFromDomainTopicStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesFromDomainTopicStatus()")
	}
	domainsTopic.Status = status

	// No error
	return nil
}

// AssignPropertiesToDomainsTopic populates the provided destination DomainsTopic from our DomainsTopic
func (domainsTopic *DomainsTopic) AssignPropertiesToDomainsTopic(destination *v1alpha1api20200601storage.DomainsTopic) error {

	// ObjectMeta
	destination.ObjectMeta = *domainsTopic.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20200601storage.DomainsTopics_Spec
	err := domainsTopic.Spec.AssignPropertiesToDomainsTopicsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesToDomainsTopicsSpec()")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20200601storage.DomainTopic_Status
	err = domainsTopic.Status.AssignPropertiesToDomainTopicStatus(&status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesToDomainTopicStatus()")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (domainsTopic *DomainsTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: domainsTopic.Spec.OriginalVersion(),
		Kind:    "DomainsTopic",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/resourceDefinitions/domains_topics
type DomainsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainsTopic `json:"items"`
}

type DomainTopic_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`

	//Name: Name of the resource.
	Name *string `json:"name,omitempty"`

	//ProvisioningState: Provisioning state of the domain topic.
	ProvisioningState *DomainTopicPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`

	//SystemData: The system metadata relating to Domain Topic resource.
	SystemData *SystemData_Status `json:"systemData,omitempty"`

	//Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DomainTopic_Status{}

// ConvertStatusFrom populates our DomainTopic_Status from the provided source
func (domainTopicStatus *DomainTopic_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20200601storage.DomainTopic_Status)
	if ok {
		// Populate our instance from source
		return domainTopicStatus.AssignPropertiesFromDomainTopicStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20200601storage.DomainTopic_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = domainTopicStatus.AssignPropertiesFromDomainTopicStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DomainTopic_Status
func (domainTopicStatus *DomainTopic_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20200601storage.DomainTopic_Status)
	if ok {
		// Populate destination from our instance
		return domainTopicStatus.AssignPropertiesToDomainTopicStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20200601storage.DomainTopic_Status{}
	err := domainTopicStatus.AssignPropertiesToDomainTopicStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &DomainTopic_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (domainTopicStatus *DomainTopic_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DomainTopic_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (domainTopicStatus *DomainTopic_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DomainTopic_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DomainTopic_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		domainTopicStatus.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		domainTopicStatus.Name = &name
	}

	// Set property ‘ProvisioningState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			domainTopicStatus.ProvisioningState = &provisioningState
		}
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_Status
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		domainTopicStatus.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		domainTopicStatus.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromDomainTopicStatus populates our DomainTopic_Status from the provided source DomainTopic_Status
func (domainTopicStatus *DomainTopic_Status) AssignPropertiesFromDomainTopicStatus(source *v1alpha1api20200601storage.DomainTopic_Status) error {

	// Conditions
	domainTopicStatus.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	domainTopicStatus.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	domainTopicStatus.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := DomainTopicPropertiesStatusProvisioningState(*source.ProvisioningState)
		domainTopicStatus.ProvisioningState = &provisioningState
	} else {
		domainTopicStatus.ProvisioningState = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "populating SystemData from SystemData, calling AssignPropertiesFromSystemDataStatus()")
		}
		domainTopicStatus.SystemData = &systemDatum
	} else {
		domainTopicStatus.SystemData = nil
	}

	// Type
	domainTopicStatus.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToDomainTopicStatus populates the provided destination DomainTopic_Status from our DomainTopic_Status
func (domainTopicStatus *DomainTopic_Status) AssignPropertiesToDomainTopicStatus(destination *v1alpha1api20200601storage.DomainTopic_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(domainTopicStatus.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(domainTopicStatus.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(domainTopicStatus.Name)

	// ProvisioningState
	if domainTopicStatus.ProvisioningState != nil {
		provisioningState := string(*domainTopicStatus.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// SystemData
	if domainTopicStatus.SystemData != nil {
		var systemDatum v1alpha1api20200601storage.SystemData_Status
		err := (*domainTopicStatus.SystemData).AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "populating SystemData from SystemData, calling AssignPropertiesToSystemDataStatus()")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(domainTopicStatus.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"2020-06-01"}
type DomainsTopicsSpecAPIVersion string

const DomainsTopicsSpecAPIVersion20200601 = DomainsTopicsSpecAPIVersion("2020-06-01")

type DomainsTopics_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"eventgrid.azure.com" json:"owner" kind:"Domain"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DomainsTopics_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (domainsTopicsSpec *DomainsTopics_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if domainsTopicsSpec == nil {
		return nil, nil
	}
	var result DomainsTopics_SpecARM

	// Set property ‘Location’:
	if domainsTopicsSpec.Location != nil {
		location := *domainsTopicsSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Tags’:
	if domainsTopicsSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range domainsTopicsSpec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (domainsTopicsSpec *DomainsTopics_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DomainsTopics_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (domainsTopicsSpec *DomainsTopics_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DomainsTopics_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DomainsTopics_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	domainsTopicsSpec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		domainsTopicsSpec.Location = &location
	}

	// Set property ‘Owner’:
	domainsTopicsSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		domainsTopicsSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			domainsTopicsSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DomainsTopics_Spec{}

// ConvertSpecFrom populates our DomainsTopics_Spec from the provided source
func (domainsTopicsSpec *DomainsTopics_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20200601storage.DomainsTopics_Spec)
	if ok {
		// Populate our instance from source
		return domainsTopicsSpec.AssignPropertiesFromDomainsTopicsSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20200601storage.DomainsTopics_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = domainsTopicsSpec.AssignPropertiesFromDomainsTopicsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DomainsTopics_Spec
func (domainsTopicsSpec *DomainsTopics_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20200601storage.DomainsTopics_Spec)
	if ok {
		// Populate destination from our instance
		return domainsTopicsSpec.AssignPropertiesToDomainsTopicsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20200601storage.DomainsTopics_Spec{}
	err := domainsTopicsSpec.AssignPropertiesToDomainsTopicsSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromDomainsTopicsSpec populates our DomainsTopics_Spec from the provided source DomainsTopics_Spec
func (domainsTopicsSpec *DomainsTopics_Spec) AssignPropertiesFromDomainsTopicsSpec(source *v1alpha1api20200601storage.DomainsTopics_Spec) error {

	// AzureName
	domainsTopicsSpec.AzureName = source.AzureName

	// Location
	domainsTopicsSpec.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	domainsTopicsSpec.Owner = source.Owner.Copy()

	// Tags
	domainsTopicsSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDomainsTopicsSpec populates the provided destination DomainsTopics_Spec from our DomainsTopics_Spec
func (domainsTopicsSpec *DomainsTopics_Spec) AssignPropertiesToDomainsTopicsSpec(destination *v1alpha1api20200601storage.DomainsTopics_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = domainsTopicsSpec.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(domainsTopicsSpec.Location)

	// OriginalVersion
	destination.OriginalVersion = domainsTopicsSpec.OriginalVersion()

	// Owner
	destination.Owner = domainsTopicsSpec.Owner.Copy()

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(domainsTopicsSpec.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (domainsTopicsSpec *DomainsTopics_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (domainsTopicsSpec *DomainsTopics_Spec) SetAzureName(azureName string) {
	domainsTopicsSpec.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&DomainsTopic{}, &DomainsTopicList{})
}
