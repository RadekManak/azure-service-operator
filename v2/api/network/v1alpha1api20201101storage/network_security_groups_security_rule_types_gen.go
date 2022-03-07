// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=networksecuritygroupssecurityrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={networksecuritygroupssecurityrules/status,networksecuritygroupssecurityrules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201101.NetworkSecurityGroupsSecurityRule
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkSecurityGroups_securityRules
type NetworkSecurityGroupsSecurityRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityGroupsSecurityRules_Spec                                   `json:"spec,omitempty"`
	Status            SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NetworkSecurityGroupsSecurityRule{}

// GetConditions returns the conditions of the resource
func (rule *NetworkSecurityGroupsSecurityRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *NetworkSecurityGroupsSecurityRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NetworkSecurityGroupsSecurityRule{}

// AzureName returns the Azure name of the resource
func (rule *NetworkSecurityGroupsSecurityRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (rule NetworkSecurityGroupsSecurityRule) GetAPIVersion() string {
	return "2020-11-01"
}

// GetResourceKind returns the kind of the resource
func (rule *NetworkSecurityGroupsSecurityRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (rule *NetworkSecurityGroupsSecurityRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *NetworkSecurityGroupsSecurityRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkSecurityGroups/securityRules"
func (rule *NetworkSecurityGroupsSecurityRule) GetType() string {
	return "Microsoft.Network/networkSecurityGroups/securityRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NetworkSecurityGroupsSecurityRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *NetworkSecurityGroupsSecurityRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *NetworkSecurityGroupsSecurityRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// Hub marks that this NetworkSecurityGroupsSecurityRule is the hub type for conversion
func (rule *NetworkSecurityGroupsSecurityRule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *NetworkSecurityGroupsSecurityRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion,
		Kind:    "NetworkSecurityGroupsSecurityRule",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.NetworkSecurityGroupsSecurityRule
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkSecurityGroups_securityRules
type NetworkSecurityGroupsSecurityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityGroupsSecurityRule `json:"items"`
}

//Storage version of v1alpha1api20201101.NetworkSecurityGroupsSecurityRules_Spec
type NetworkSecurityGroupsSecurityRules_Spec struct {
	Access *string `json:"access,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName                            string        `json:"azureName"`
	Description                          *string       `json:"description,omitempty"`
	DestinationAddressPrefix             *string       `json:"destinationAddressPrefix,omitempty"`
	DestinationAddressPrefixes           []string      `json:"destinationAddressPrefixes,omitempty"`
	DestinationApplicationSecurityGroups []SubResource `json:"destinationApplicationSecurityGroups,omitempty"`
	DestinationPortRange                 *string       `json:"destinationPortRange,omitempty"`
	DestinationPortRanges                []string      `json:"destinationPortRanges,omitempty"`
	Direction                            *string       `json:"direction,omitempty"`
	Location                             *string       `json:"location,omitempty"`
	OriginalVersion                      string        `json:"originalVersion"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a network.azure.com/NetworkSecurityGroup resource
	Owner                           genruntime.KnownResourceReference `group:"network.azure.com" json:"owner" kind:"NetworkSecurityGroup"`
	Priority                        *int                              `json:"priority,omitempty"`
	PropertyBag                     genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Protocol                        *string                           `json:"protocol,omitempty"`
	SourceAddressPrefix             *string                           `json:"sourceAddressPrefix,omitempty"`
	SourceAddressPrefixes           []string                          `json:"sourceAddressPrefixes,omitempty"`
	SourceApplicationSecurityGroups []SubResource                     `json:"sourceApplicationSecurityGroups,omitempty"`
	SourcePortRange                 *string                           `json:"sourcePortRange,omitempty"`
	SourcePortRanges                []string                          `json:"sourcePortRanges,omitempty"`
	Tags                            map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NetworkSecurityGroupsSecurityRules_Spec{}

// ConvertSpecFrom populates our NetworkSecurityGroupsSecurityRules_Spec from the provided source
func (rules *NetworkSecurityGroupsSecurityRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == rules {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(rules)
}

// ConvertSpecTo populates the provided destination from our NetworkSecurityGroupsSecurityRules_Spec
func (rules *NetworkSecurityGroupsSecurityRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == rules {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(rules)
}

//Storage version of v1alpha1api20201101.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
type SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded struct {
	Access                               *string                                                                                 `json:"access,omitempty"`
	Conditions                           []conditions.Condition                                                                  `json:"conditions,omitempty"`
	Description                          *string                                                                                 `json:"description,omitempty"`
	DestinationAddressPrefix             *string                                                                                 `json:"destinationAddressPrefix,omitempty"`
	DestinationAddressPrefixes           []string                                                                                `json:"destinationAddressPrefixes,omitempty"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded `json:"destinationApplicationSecurityGroups,omitempty"`
	DestinationPortRange                 *string                                                                                 `json:"destinationPortRange,omitempty"`
	DestinationPortRanges                []string                                                                                `json:"destinationPortRanges,omitempty"`
	Direction                            *string                                                                                 `json:"direction,omitempty"`
	Etag                                 *string                                                                                 `json:"etag,omitempty"`
	Id                                   *string                                                                                 `json:"id,omitempty"`
	Name                                 *string                                                                                 `json:"name,omitempty"`
	Priority                             *int                                                                                    `json:"priority,omitempty"`
	PropertyBag                          genruntime.PropertyBag                                                                  `json:"$propertyBag,omitempty"`
	Protocol                             *string                                                                                 `json:"protocol,omitempty"`
	ProvisioningState                    *string                                                                                 `json:"provisioningState,omitempty"`
	SourceAddressPrefix                  *string                                                                                 `json:"sourceAddressPrefix,omitempty"`
	SourceAddressPrefixes                []string                                                                                `json:"sourceAddressPrefixes,omitempty"`
	SourceApplicationSecurityGroups      []ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded `json:"sourceApplicationSecurityGroups,omitempty"`
	SourcePortRange                      *string                                                                                 `json:"sourcePortRange,omitempty"`
	SourcePortRanges                     []string                                                                                `json:"sourcePortRanges,omitempty"`
	Type                                 *string                                                                                 `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}

// ConvertStatusFrom populates our SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded from the provided source
func (embedded *SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == embedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(embedded)
}

// ConvertStatusTo populates the provided destination from our SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
func (embedded *SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == embedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(embedded)
}

//Storage version of v1alpha1api20201101.ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
type ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityGroupsSecurityRule{}, &NetworkSecurityGroupsSecurityRuleList{})
}
