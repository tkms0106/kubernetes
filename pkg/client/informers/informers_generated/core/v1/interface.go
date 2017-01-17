/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by informer-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1alpha1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1] --internal-clientset-package=k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset --listers-package=k8s.io/kubernetes/pkg/client/listers --versioned-clientset-package=k8s.io/kubernetes/pkg/client/clientset_generated/clientset

package v1

import (
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ComponentStatuses returns a ComponentStatusInformer.
	ComponentStatuses() ComponentStatusInformer
	// ConfigMaps returns a ConfigMapInformer.
	ConfigMaps() ConfigMapInformer
	// Endpoints returns a EndpointsInformer.
	Endpoints() EndpointsInformer
	// Events returns a EventInformer.
	Events() EventInformer
	// LimitRanges returns a LimitRangeInformer.
	LimitRanges() LimitRangeInformer
	// Namespaces returns a NamespaceInformer.
	Namespaces() NamespaceInformer
	// Nodes returns a NodeInformer.
	Nodes() NodeInformer
	// PersistentVolumes returns a PersistentVolumeInformer.
	PersistentVolumes() PersistentVolumeInformer
	// PersistentVolumeClaims returns a PersistentVolumeClaimInformer.
	PersistentVolumeClaims() PersistentVolumeClaimInformer
	// Pods returns a PodInformer.
	Pods() PodInformer
	// PodTemplates returns a PodTemplateInformer.
	PodTemplates() PodTemplateInformer
	// ReplicationControllers returns a ReplicationControllerInformer.
	ReplicationControllers() ReplicationControllerInformer
	// ResourceQuotas returns a ResourceQuotaInformer.
	ResourceQuotas() ResourceQuotaInformer
	// Secrets returns a SecretInformer.
	Secrets() SecretInformer
	// Services returns a ServiceInformer.
	Services() ServiceInformer
	// ServiceAccounts returns a ServiceAccountInformer.
	ServiceAccounts() ServiceAccountInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// ComponentStatuses returns a ComponentStatusInformer.
func (v *version) ComponentStatuses() ComponentStatusInformer {
	return &componentStatusInformer{factory: v.SharedInformerFactory}
}

// ConfigMaps returns a ConfigMapInformer.
func (v *version) ConfigMaps() ConfigMapInformer {
	return &configMapInformer{factory: v.SharedInformerFactory}
}

// Endpoints returns a EndpointsInformer.
func (v *version) Endpoints() EndpointsInformer {
	return &endpointsInformer{factory: v.SharedInformerFactory}
}

// Events returns a EventInformer.
func (v *version) Events() EventInformer {
	return &eventInformer{factory: v.SharedInformerFactory}
}

// LimitRanges returns a LimitRangeInformer.
func (v *version) LimitRanges() LimitRangeInformer {
	return &limitRangeInformer{factory: v.SharedInformerFactory}
}

// Namespaces returns a NamespaceInformer.
func (v *version) Namespaces() NamespaceInformer {
	return &namespaceInformer{factory: v.SharedInformerFactory}
}

// Nodes returns a NodeInformer.
func (v *version) Nodes() NodeInformer {
	return &nodeInformer{factory: v.SharedInformerFactory}
}

// PersistentVolumes returns a PersistentVolumeInformer.
func (v *version) PersistentVolumes() PersistentVolumeInformer {
	return &persistentVolumeInformer{factory: v.SharedInformerFactory}
}

// PersistentVolumeClaims returns a PersistentVolumeClaimInformer.
func (v *version) PersistentVolumeClaims() PersistentVolumeClaimInformer {
	return &persistentVolumeClaimInformer{factory: v.SharedInformerFactory}
}

// Pods returns a PodInformer.
func (v *version) Pods() PodInformer {
	return &podInformer{factory: v.SharedInformerFactory}
}

// PodTemplates returns a PodTemplateInformer.
func (v *version) PodTemplates() PodTemplateInformer {
	return &podTemplateInformer{factory: v.SharedInformerFactory}
}

// ReplicationControllers returns a ReplicationControllerInformer.
func (v *version) ReplicationControllers() ReplicationControllerInformer {
	return &replicationControllerInformer{factory: v.SharedInformerFactory}
}

// ResourceQuotas returns a ResourceQuotaInformer.
func (v *version) ResourceQuotas() ResourceQuotaInformer {
	return &resourceQuotaInformer{factory: v.SharedInformerFactory}
}

// Secrets returns a SecretInformer.
func (v *version) Secrets() SecretInformer {
	return &secretInformer{factory: v.SharedInformerFactory}
}

// Services returns a ServiceInformer.
func (v *version) Services() ServiceInformer {
	return &serviceInformer{factory: v.SharedInformerFactory}
}

// ServiceAccounts returns a ServiceAccountInformer.
func (v *version) ServiceAccounts() ServiceAccountInformer {
	return &serviceAccountInformer{factory: v.SharedInformerFactory}
}
