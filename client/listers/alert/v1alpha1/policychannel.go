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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-newrelic-api/apis/alert/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PolicyChannelLister helps list PolicyChannels.
// All objects returned here must be treated as read-only.
type PolicyChannelLister interface {
	// List lists all PolicyChannels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PolicyChannel, err error)
	// PolicyChannels returns an object that can list and get PolicyChannels.
	PolicyChannels(namespace string) PolicyChannelNamespaceLister
	PolicyChannelListerExpansion
}

// policyChannelLister implements the PolicyChannelLister interface.
type policyChannelLister struct {
	indexer cache.Indexer
}

// NewPolicyChannelLister returns a new PolicyChannelLister.
func NewPolicyChannelLister(indexer cache.Indexer) PolicyChannelLister {
	return &policyChannelLister{indexer: indexer}
}

// List lists all PolicyChannels in the indexer.
func (s *policyChannelLister) List(selector labels.Selector) (ret []*v1alpha1.PolicyChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PolicyChannel))
	})
	return ret, err
}

// PolicyChannels returns an object that can list and get PolicyChannels.
func (s *policyChannelLister) PolicyChannels(namespace string) PolicyChannelNamespaceLister {
	return policyChannelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PolicyChannelNamespaceLister helps list and get PolicyChannels.
// All objects returned here must be treated as read-only.
type PolicyChannelNamespaceLister interface {
	// List lists all PolicyChannels in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PolicyChannel, err error)
	// Get retrieves the PolicyChannel from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PolicyChannel, error)
	PolicyChannelNamespaceListerExpansion
}

// policyChannelNamespaceLister implements the PolicyChannelNamespaceLister
// interface.
type policyChannelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PolicyChannels in the indexer for a given namespace.
func (s policyChannelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PolicyChannel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PolicyChannel))
	})
	return ret, err
}

// Get retrieves the PolicyChannel from the indexer for a given namespace and name.
func (s policyChannelNamespaceLister) Get(name string) (*v1alpha1.PolicyChannel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("policychannel"), name)
	}
	return obj.(*v1alpha1.PolicyChannel), nil
}
