// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101previewstorage

import (
	"fmt"
	v20210101ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1beta20210101previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210101preview.NamespacesTopic
//Deprecated version of NamespacesTopic. Use v1beta20210101preview.NamespacesTopic instead
type NamespacesTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesTopics_Spec `json:"spec,omitempty"`
	Status            SBTopic_Status        `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesTopic{}

// GetConditions returns the conditions of the resource
func (topic *NamespacesTopic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *NamespacesTopic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesTopic{}

// ConvertFrom populates our NamespacesTopic from the provided hub NamespacesTopic
func (topic *NamespacesTopic) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210101ps.NamespacesTopic)
	if !ok {
		return fmt.Errorf("expected servicebus/v1beta20210101previewstorage/NamespacesTopic but received %T instead", hub)
	}

	return topic.AssignPropertiesFromNamespacesTopic(source)
}

// ConvertTo populates the provided hub NamespacesTopic from our NamespacesTopic
func (topic *NamespacesTopic) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210101ps.NamespacesTopic)
	if !ok {
		return fmt.Errorf("expected servicebus/v1beta20210101previewstorage/NamespacesTopic but received %T instead", hub)
	}

	return topic.AssignPropertiesToNamespacesTopic(destination)
}

var _ genruntime.KubernetesResource = &NamespacesTopic{}

// AzureName returns the Azure name of the resource
func (topic *NamespacesTopic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (topic NamespacesTopic) GetAPIVersion() string {
	return "2021-01-01-preview"
}

// GetResourceKind returns the kind of the resource
func (topic *NamespacesTopic) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (topic *NamespacesTopic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *NamespacesTopic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics"
func (topic *NamespacesTopic) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *NamespacesTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SBTopic_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (topic *NamespacesTopic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  topic.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (topic *NamespacesTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SBTopic_Status); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st SBTopic_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// AssignPropertiesFromNamespacesTopic populates our NamespacesTopic from the provided source NamespacesTopic
func (topic *NamespacesTopic) AssignPropertiesFromNamespacesTopic(source *v20210101ps.NamespacesTopic) error {

	// ObjectMeta
	topic.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesTopics_Spec
	err := spec.AssignPropertiesFromNamespacesTopicsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNamespacesTopicsSpec() to populate field Spec")
	}
	topic.Spec = spec

	// Status
	var status SBTopic_Status
	err = status.AssignPropertiesFromSBTopicStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSBTopicStatus() to populate field Status")
	}
	topic.Status = status

	// No error
	return nil
}

// AssignPropertiesToNamespacesTopic populates the provided destination NamespacesTopic from our NamespacesTopic
func (topic *NamespacesTopic) AssignPropertiesToNamespacesTopic(destination *v20210101ps.NamespacesTopic) error {

	// ObjectMeta
	destination.ObjectMeta = *topic.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210101ps.NamespacesTopics_Spec
	err := topic.Spec.AssignPropertiesToNamespacesTopicsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNamespacesTopicsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210101ps.SBTopic_Status
	err = topic.Status.AssignPropertiesToSBTopicStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSBTopicStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *NamespacesTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion,
		Kind:    "NamespacesTopic",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210101preview.NamespacesTopic
//Deprecated version of NamespacesTopic. Use v1beta20210101preview.NamespacesTopic instead
type NamespacesTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesTopic `json:"items"`
}

//Storage version of v1alpha1api20210101preview.NamespacesTopics_Spec
type NamespacesTopics_Spec struct {
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// +kubebuilder:validation:MinLength=1
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName                           string  `json:"azureName,omitempty"`
	DefaultMessageTimeToLive            *string `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool   `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool   `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool   `json:"enablePartitioning,omitempty"`
	Location                            *string `json:"location,omitempty"`
	MaxSizeInMegabytes                  *int    `json:"maxSizeInMegabytes,omitempty"`
	OriginalVersion                     string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a servicebus.azure.com/Namespace resource
	Owner                      *genruntime.KnownResourceReference `group:"servicebus.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PropertyBag                genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection *bool                              `json:"requiresDuplicateDetection,omitempty"`
	SupportOrdering            *bool                              `json:"supportOrdering,omitempty"`
	Tags                       map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesTopics_Spec{}

// ConvertSpecFrom populates our NamespacesTopics_Spec from the provided source
func (topics *NamespacesTopics_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210101ps.NamespacesTopics_Spec)
	if ok {
		// Populate our instance from source
		return topics.AssignPropertiesFromNamespacesTopicsSpec(src)
	}

	// Convert to an intermediate form
	src = &v20210101ps.NamespacesTopics_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = topics.AssignPropertiesFromNamespacesTopicsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesTopics_Spec
func (topics *NamespacesTopics_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210101ps.NamespacesTopics_Spec)
	if ok {
		// Populate destination from our instance
		return topics.AssignPropertiesToNamespacesTopicsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210101ps.NamespacesTopics_Spec{}
	err := topics.AssignPropertiesToNamespacesTopicsSpec(dst)
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

// AssignPropertiesFromNamespacesTopicsSpec populates our NamespacesTopics_Spec from the provided source NamespacesTopics_Spec
func (topics *NamespacesTopics_Spec) AssignPropertiesFromNamespacesTopicsSpec(source *v20210101ps.NamespacesTopics_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AutoDeleteOnIdle
	topics.AutoDeleteOnIdle = genruntime.ClonePointerToString(source.AutoDeleteOnIdle)

	// AzureName
	topics.AzureName = source.AzureName

	// DefaultMessageTimeToLive
	topics.DefaultMessageTimeToLive = genruntime.ClonePointerToString(source.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	topics.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(source.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if source.EnableBatchedOperations != nil {
		enableBatchedOperation := *source.EnableBatchedOperations
		topics.EnableBatchedOperations = &enableBatchedOperation
	} else {
		topics.EnableBatchedOperations = nil
	}

	// EnableExpress
	if source.EnableExpress != nil {
		enableExpress := *source.EnableExpress
		topics.EnableExpress = &enableExpress
	} else {
		topics.EnableExpress = nil
	}

	// EnablePartitioning
	if source.EnablePartitioning != nil {
		enablePartitioning := *source.EnablePartitioning
		topics.EnablePartitioning = &enablePartitioning
	} else {
		topics.EnablePartitioning = nil
	}

	// Location
	topics.Location = genruntime.ClonePointerToString(source.Location)

	// MaxSizeInMegabytes
	topics.MaxSizeInMegabytes = genruntime.ClonePointerToInt(source.MaxSizeInMegabytes)

	// OriginalVersion
	topics.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		topics.Owner = &owner
	} else {
		topics.Owner = nil
	}

	// RequiresDuplicateDetection
	if source.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *source.RequiresDuplicateDetection
		topics.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		topics.RequiresDuplicateDetection = nil
	}

	// SupportOrdering
	if source.SupportOrdering != nil {
		supportOrdering := *source.SupportOrdering
		topics.SupportOrdering = &supportOrdering
	} else {
		topics.SupportOrdering = nil
	}

	// Tags
	topics.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		topics.PropertyBag = propertyBag
	} else {
		topics.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNamespacesTopicsSpec populates the provided destination NamespacesTopics_Spec from our NamespacesTopics_Spec
func (topics *NamespacesTopics_Spec) AssignPropertiesToNamespacesTopicsSpec(destination *v20210101ps.NamespacesTopics_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(topics.PropertyBag)

	// AutoDeleteOnIdle
	destination.AutoDeleteOnIdle = genruntime.ClonePointerToString(topics.AutoDeleteOnIdle)

	// AzureName
	destination.AzureName = topics.AzureName

	// DefaultMessageTimeToLive
	destination.DefaultMessageTimeToLive = genruntime.ClonePointerToString(topics.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	destination.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(topics.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if topics.EnableBatchedOperations != nil {
		enableBatchedOperation := *topics.EnableBatchedOperations
		destination.EnableBatchedOperations = &enableBatchedOperation
	} else {
		destination.EnableBatchedOperations = nil
	}

	// EnableExpress
	if topics.EnableExpress != nil {
		enableExpress := *topics.EnableExpress
		destination.EnableExpress = &enableExpress
	} else {
		destination.EnableExpress = nil
	}

	// EnablePartitioning
	if topics.EnablePartitioning != nil {
		enablePartitioning := *topics.EnablePartitioning
		destination.EnablePartitioning = &enablePartitioning
	} else {
		destination.EnablePartitioning = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(topics.Location)

	// MaxSizeInMegabytes
	destination.MaxSizeInMegabytes = genruntime.ClonePointerToInt(topics.MaxSizeInMegabytes)

	// OriginalVersion
	destination.OriginalVersion = topics.OriginalVersion

	// Owner
	if topics.Owner != nil {
		owner := topics.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// RequiresDuplicateDetection
	if topics.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *topics.RequiresDuplicateDetection
		destination.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		destination.RequiresDuplicateDetection = nil
	}

	// SupportOrdering
	if topics.SupportOrdering != nil {
		supportOrdering := *topics.SupportOrdering
		destination.SupportOrdering = &supportOrdering
	} else {
		destination.SupportOrdering = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(topics.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20210101preview.SBTopic_Status
//Deprecated version of SBTopic_Status. Use v1beta20210101preview.SBTopic_Status instead
type SBTopic_Status struct {
	AccessedAt                          *string                     `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                     `json:"autoDeleteOnIdle,omitempty"`
	Conditions                          []conditions.Condition      `json:"conditions,omitempty"`
	CountDetails                        *MessageCountDetails_Status `json:"countDetails,omitempty"`
	CreatedAt                           *string                     `json:"createdAt,omitempty"`
	DefaultMessageTimeToLive            *string                     `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                     `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                       `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                       `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                       `json:"enablePartitioning,omitempty"`
	Id                                  *string                     `json:"id,omitempty"`
	MaxSizeInMegabytes                  *int                        `json:"maxSizeInMegabytes,omitempty"`
	Name                                *string                     `json:"name,omitempty"`
	PropertyBag                         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection          *bool                       `json:"requiresDuplicateDetection,omitempty"`
	SizeInBytes                         *int                        `json:"sizeInBytes,omitempty"`
	Status                              *string                     `json:"status,omitempty"`
	SubscriptionCount                   *int                        `json:"subscriptionCount,omitempty"`
	SupportOrdering                     *bool                       `json:"supportOrdering,omitempty"`
	SystemData                          *SystemData_Status          `json:"systemData,omitempty"`
	Type                                *string                     `json:"type,omitempty"`
	UpdatedAt                           *string                     `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SBTopic_Status{}

// ConvertStatusFrom populates our SBTopic_Status from the provided source
func (topic *SBTopic_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210101ps.SBTopic_Status)
	if ok {
		// Populate our instance from source
		return topic.AssignPropertiesFromSBTopicStatus(src)
	}

	// Convert to an intermediate form
	src = &v20210101ps.SBTopic_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = topic.AssignPropertiesFromSBTopicStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SBTopic_Status
func (topic *SBTopic_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210101ps.SBTopic_Status)
	if ok {
		// Populate destination from our instance
		return topic.AssignPropertiesToSBTopicStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20210101ps.SBTopic_Status{}
	err := topic.AssignPropertiesToSBTopicStatus(dst)
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

// AssignPropertiesFromSBTopicStatus populates our SBTopic_Status from the provided source SBTopic_Status
func (topic *SBTopic_Status) AssignPropertiesFromSBTopicStatus(source *v20210101ps.SBTopic_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AccessedAt
	topic.AccessedAt = genruntime.ClonePointerToString(source.AccessedAt)

	// AutoDeleteOnIdle
	topic.AutoDeleteOnIdle = genruntime.ClonePointerToString(source.AutoDeleteOnIdle)

	// Conditions
	topic.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CountDetails
	if source.CountDetails != nil {
		var countDetail MessageCountDetails_Status
		err := countDetail.AssignPropertiesFromMessageCountDetailsStatus(source.CountDetails)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromMessageCountDetailsStatus() to populate field CountDetails")
		}
		topic.CountDetails = &countDetail
	} else {
		topic.CountDetails = nil
	}

	// CreatedAt
	topic.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// DefaultMessageTimeToLive
	topic.DefaultMessageTimeToLive = genruntime.ClonePointerToString(source.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	topic.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(source.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if source.EnableBatchedOperations != nil {
		enableBatchedOperation := *source.EnableBatchedOperations
		topic.EnableBatchedOperations = &enableBatchedOperation
	} else {
		topic.EnableBatchedOperations = nil
	}

	// EnableExpress
	if source.EnableExpress != nil {
		enableExpress := *source.EnableExpress
		topic.EnableExpress = &enableExpress
	} else {
		topic.EnableExpress = nil
	}

	// EnablePartitioning
	if source.EnablePartitioning != nil {
		enablePartitioning := *source.EnablePartitioning
		topic.EnablePartitioning = &enablePartitioning
	} else {
		topic.EnablePartitioning = nil
	}

	// Id
	topic.Id = genruntime.ClonePointerToString(source.Id)

	// MaxSizeInMegabytes
	topic.MaxSizeInMegabytes = genruntime.ClonePointerToInt(source.MaxSizeInMegabytes)

	// Name
	topic.Name = genruntime.ClonePointerToString(source.Name)

	// RequiresDuplicateDetection
	if source.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *source.RequiresDuplicateDetection
		topic.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		topic.RequiresDuplicateDetection = nil
	}

	// SizeInBytes
	topic.SizeInBytes = genruntime.ClonePointerToInt(source.SizeInBytes)

	// Status
	topic.Status = genruntime.ClonePointerToString(source.Status)

	// SubscriptionCount
	topic.SubscriptionCount = genruntime.ClonePointerToInt(source.SubscriptionCount)

	// SupportOrdering
	if source.SupportOrdering != nil {
		supportOrdering := *source.SupportOrdering
		topic.SupportOrdering = &supportOrdering
	} else {
		topic.SupportOrdering = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
		}
		topic.SystemData = &systemDatum
	} else {
		topic.SystemData = nil
	}

	// Type
	topic.Type = genruntime.ClonePointerToString(source.Type)

	// UpdatedAt
	topic.UpdatedAt = genruntime.ClonePointerToString(source.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		topic.PropertyBag = propertyBag
	} else {
		topic.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSBTopicStatus populates the provided destination SBTopic_Status from our SBTopic_Status
func (topic *SBTopic_Status) AssignPropertiesToSBTopicStatus(destination *v20210101ps.SBTopic_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(topic.PropertyBag)

	// AccessedAt
	destination.AccessedAt = genruntime.ClonePointerToString(topic.AccessedAt)

	// AutoDeleteOnIdle
	destination.AutoDeleteOnIdle = genruntime.ClonePointerToString(topic.AutoDeleteOnIdle)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(topic.Conditions)

	// CountDetails
	if topic.CountDetails != nil {
		var countDetail v20210101ps.MessageCountDetails_Status
		err := topic.CountDetails.AssignPropertiesToMessageCountDetailsStatus(&countDetail)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToMessageCountDetailsStatus() to populate field CountDetails")
		}
		destination.CountDetails = &countDetail
	} else {
		destination.CountDetails = nil
	}

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(topic.CreatedAt)

	// DefaultMessageTimeToLive
	destination.DefaultMessageTimeToLive = genruntime.ClonePointerToString(topic.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	destination.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(topic.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if topic.EnableBatchedOperations != nil {
		enableBatchedOperation := *topic.EnableBatchedOperations
		destination.EnableBatchedOperations = &enableBatchedOperation
	} else {
		destination.EnableBatchedOperations = nil
	}

	// EnableExpress
	if topic.EnableExpress != nil {
		enableExpress := *topic.EnableExpress
		destination.EnableExpress = &enableExpress
	} else {
		destination.EnableExpress = nil
	}

	// EnablePartitioning
	if topic.EnablePartitioning != nil {
		enablePartitioning := *topic.EnablePartitioning
		destination.EnablePartitioning = &enablePartitioning
	} else {
		destination.EnablePartitioning = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(topic.Id)

	// MaxSizeInMegabytes
	destination.MaxSizeInMegabytes = genruntime.ClonePointerToInt(topic.MaxSizeInMegabytes)

	// Name
	destination.Name = genruntime.ClonePointerToString(topic.Name)

	// RequiresDuplicateDetection
	if topic.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *topic.RequiresDuplicateDetection
		destination.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		destination.RequiresDuplicateDetection = nil
	}

	// SizeInBytes
	destination.SizeInBytes = genruntime.ClonePointerToInt(topic.SizeInBytes)

	// Status
	destination.Status = genruntime.ClonePointerToString(topic.Status)

	// SubscriptionCount
	destination.SubscriptionCount = genruntime.ClonePointerToInt(topic.SubscriptionCount)

	// SupportOrdering
	if topic.SupportOrdering != nil {
		supportOrdering := *topic.SupportOrdering
		destination.SupportOrdering = &supportOrdering
	} else {
		destination.SupportOrdering = nil
	}

	// SystemData
	if topic.SystemData != nil {
		var systemDatum v20210101ps.SystemData_Status
		err := topic.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(topic.Type)

	// UpdatedAt
	destination.UpdatedAt = genruntime.ClonePointerToString(topic.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&NamespacesTopic{}, &NamespacesTopicList{})
}
