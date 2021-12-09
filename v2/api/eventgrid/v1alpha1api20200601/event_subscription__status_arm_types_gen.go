// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

type EventSubscription_StatusARM struct {
	//Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`

	//Name: Name of the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the event subscription.
	Properties *EventSubscriptionProperties_StatusARM `json:"properties,omitempty"`

	//SystemData: The system metadata relating to Event Subscription resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	//Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

type EventSubscriptionProperties_StatusARM struct {
	//DeadLetterDestination: The DeadLetter destination of the event subscription.
	DeadLetterDestination *DeadLetterDestination_StatusARM `json:"deadLetterDestination,omitempty"`

	//Destination: Information about the destination where events have to be delivered
	//for the event subscription.
	Destination *EventSubscriptionDestination_StatusARM `json:"destination,omitempty"`

	//EventDeliverySchema: The event delivery schema for the event subscription.
	EventDeliverySchema *EventSubscriptionPropertiesStatusEventDeliverySchema `json:"eventDeliverySchema,omitempty"`

	//ExpirationTimeUtc: Expiration time of the event subscription.
	ExpirationTimeUtc *string `json:"expirationTimeUtc,omitempty"`

	//Filter: Information about the filter for the event subscription.
	Filter *EventSubscriptionFilter_StatusARM `json:"filter,omitempty"`

	//Labels: List of user defined labels.
	Labels []string `json:"labels,omitempty"`

	//ProvisioningState: Provisioning state of the event subscription.
	ProvisioningState *EventSubscriptionPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`

	//RetryPolicy: The retry policy for events. This can be used to configure maximum
	//number of delivery attempts and time to live for events.
	RetryPolicy *RetryPolicy_StatusARM `json:"retryPolicy,omitempty"`

	//Topic: Name of the topic of the event subscription.
	Topic *string `json:"topic,omitempty"`
}

type DeadLetterDestination_StatusARM struct {
	//EndpointType: Type of the endpoint for the dead letter destination
	EndpointType DeadLetterDestinationStatusEndpointType `json:"endpointType"`
}

type EventSubscriptionDestination_StatusARM struct {
	//EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType EventSubscriptionDestinationStatusEndpointType `json:"endpointType"`
}

type EventSubscriptionFilter_StatusARM struct {
	//AdvancedFilters: An array of advanced filters that are used for filtering event
	//subscriptions.
	AdvancedFilters []AdvancedFilter_StatusARM `json:"advancedFilters,omitempty"`

	//IncludedEventTypes: A list of applicable event types that need to be part of the
	//event subscription. If it is desired to subscribe to all default event types,
	//set the IncludedEventTypes to null.
	IncludedEventTypes []string `json:"includedEventTypes,omitempty"`

	//IsSubjectCaseSensitive: Specifies if the SubjectBeginsWith and SubjectEndsWith
	//properties of the filter
	//should be compared in a case sensitive manner.
	IsSubjectCaseSensitive *bool `json:"isSubjectCaseSensitive,omitempty"`

	//SubjectBeginsWith: An optional string to filter events for an event subscription
	//based on a resource path prefix.
	//The format of this depends on the publisher of the events.
	//Wildcard characters are not supported in this path.
	SubjectBeginsWith *string `json:"subjectBeginsWith,omitempty"`

	//SubjectEndsWith: An optional string to filter events for an event subscription
	//based on a resource path suffix.
	//Wildcard characters are not supported in this path.
	SubjectEndsWith *string `json:"subjectEndsWith,omitempty"`
}

type RetryPolicy_StatusARM struct {
	//EventTimeToLiveInMinutes: Time To Live (in minutes) for events.
	EventTimeToLiveInMinutes *int `json:"eventTimeToLiveInMinutes,omitempty"`

	//MaxDeliveryAttempts: Maximum number of delivery retry attempts for events.
	MaxDeliveryAttempts *int `json:"maxDeliveryAttempts,omitempty"`
}

type AdvancedFilter_StatusARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	//OperatorType: The operator type used for filtering, e.g., NumberIn,
	//StringContains, BoolEquals and others.
	OperatorType AdvancedFilterStatusOperatorType `json:"operatorType"`
}
