// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	alpha20210501 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1alpha1api20210501"
	alpha20210501s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1alpha1api20210501storage"
	v20210501 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1beta20210501"
	v20210501s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1beta20210501storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ManagedClusterExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ManagedClusterExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&alpha20210501.ManagedCluster{},
		&alpha20210501s.ManagedCluster{},
		&v20210501.ManagedCluster{},
		&v20210501s.ManagedCluster{}}
}
