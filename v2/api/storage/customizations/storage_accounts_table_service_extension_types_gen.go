// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20220901 "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901"
	v20220901s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type StorageAccountsTableServiceExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *StorageAccountsTableServiceExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20220901.StorageAccountsTableService{},
		&v20220901s.StorageAccountsTableService{}}
}
