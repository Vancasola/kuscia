// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/secretflow/kuscia/pkg/crd/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AppImages returns a AppImageInformer.
	AppImages() AppImageInformer
	// ClusterDomainRoutes returns a ClusterDomainRouteInformer.
	ClusterDomainRoutes() ClusterDomainRouteInformer
	// Domains returns a DomainInformer.
	Domains() DomainInformer
	// DomainAppImages returns a DomainAppImageInformer.
	DomainAppImages() DomainAppImageInformer
	// DomainDatas returns a DomainDataInformer.
	DomainDatas() DomainDataInformer
	// DomainDataGrants returns a DomainDataGrantInformer.
	DomainDataGrants() DomainDataGrantInformer
	// DomainDataSources returns a DomainDataSourceInformer.
	DomainDataSources() DomainDataSourceInformer
	// DomainRoutes returns a DomainRouteInformer.
	DomainRoutes() DomainRouteInformer
	// Gateways returns a GatewayInformer.
	Gateways() GatewayInformer
	// InteropConfigs returns a InteropConfigInformer.
	InteropConfigs() InteropConfigInformer
	// KusciaDeployments returns a KusciaDeploymentInformer.
	KusciaDeployments() KusciaDeploymentInformer
	// KusciaDeploymentSummaries returns a KusciaDeploymentSummaryInformer.
	KusciaDeploymentSummaries() KusciaDeploymentSummaryInformer
	// KusciaJobs returns a KusciaJobInformer.
	KusciaJobs() KusciaJobInformer
	// KusciaJobSummaries returns a KusciaJobSummaryInformer.
	KusciaJobSummaries() KusciaJobSummaryInformer
	// KusciaTasks returns a KusciaTaskInformer.
	KusciaTasks() KusciaTaskInformer
	// KusciaTaskSummaries returns a KusciaTaskSummaryInformer.
	KusciaTaskSummaries() KusciaTaskSummaryInformer
	// TaskResources returns a TaskResourceInformer.
	TaskResources() TaskResourceInformer
	// TaskResourceGroups returns a TaskResourceGroupInformer.
	TaskResourceGroups() TaskResourceGroupInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AppImages returns a AppImageInformer.
func (v *version) AppImages() AppImageInformer {
	return &appImageInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterDomainRoutes returns a ClusterDomainRouteInformer.
func (v *version) ClusterDomainRoutes() ClusterDomainRouteInformer {
	return &clusterDomainRouteInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Domains returns a DomainInformer.
func (v *version) Domains() DomainInformer {
	return &domainInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// DomainAppImages returns a DomainAppImageInformer.
func (v *version) DomainAppImages() DomainAppImageInformer {
	return &domainAppImageInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainDatas returns a DomainDataInformer.
func (v *version) DomainDatas() DomainDataInformer {
	return &domainDataInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainDataGrants returns a DomainDataGrantInformer.
func (v *version) DomainDataGrants() DomainDataGrantInformer {
	return &domainDataGrantInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainDataSources returns a DomainDataSourceInformer.
func (v *version) DomainDataSources() DomainDataSourceInformer {
	return &domainDataSourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainRoutes returns a DomainRouteInformer.
func (v *version) DomainRoutes() DomainRouteInformer {
	return &domainRouteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Gateways returns a GatewayInformer.
func (v *version) Gateways() GatewayInformer {
	return &gatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InteropConfigs returns a InteropConfigInformer.
func (v *version) InteropConfigs() InteropConfigInformer {
	return &interopConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// KusciaDeployments returns a KusciaDeploymentInformer.
func (v *version) KusciaDeployments() KusciaDeploymentInformer {
	return &kusciaDeploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// KusciaDeploymentSummaries returns a KusciaDeploymentSummaryInformer.
func (v *version) KusciaDeploymentSummaries() KusciaDeploymentSummaryInformer {
	return &kusciaDeploymentSummaryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// KusciaJobs returns a KusciaJobInformer.
func (v *version) KusciaJobs() KusciaJobInformer {
	return &kusciaJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// KusciaJobSummaries returns a KusciaJobSummaryInformer.
func (v *version) KusciaJobSummaries() KusciaJobSummaryInformer {
	return &kusciaJobSummaryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// KusciaTasks returns a KusciaTaskInformer.
func (v *version) KusciaTasks() KusciaTaskInformer {
	return &kusciaTaskInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// KusciaTaskSummaries returns a KusciaTaskSummaryInformer.
func (v *version) KusciaTaskSummaries() KusciaTaskSummaryInformer {
	return &kusciaTaskSummaryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TaskResources returns a TaskResourceInformer.
func (v *version) TaskResources() TaskResourceInformer {
	return &taskResourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TaskResourceGroups returns a TaskResourceGroupInformer.
func (v *version) TaskResourceGroups() TaskResourceGroupInformer {
	return &taskResourceGroupInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
