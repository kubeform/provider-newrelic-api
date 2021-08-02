/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	alertv1alpha1 "kubeform.dev/provider-newrelic-api/apis/alert/v1alpha1"
	versioned "kubeform.dev/provider-newrelic-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-newrelic-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-newrelic-api/client/listers/alert/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PolicyInformer provides access to a shared informer and lister for
// Policies.
type PolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PolicyLister
}

type policyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPolicyInformer constructs a new informer for Policy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPolicyInformer constructs a new informer for Policy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AlertV1alpha1().Policies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AlertV1alpha1().Policies(namespace).Watch(context.TODO(), options)
			},
		},
		&alertv1alpha1.Policy{},
		resyncPeriod,
		indexers,
	)
}

func (f *policyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *policyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&alertv1alpha1.Policy{}, f.defaultInformer)
}

func (f *policyInformer) Lister() v1alpha1.PolicyLister {
	return v1alpha1.NewPolicyLister(f.Informer().GetIndexer())
}
