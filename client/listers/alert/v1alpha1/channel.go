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

// ChannelLister helps list Channels.
// All objects returned here must be treated as read-only.
type ChannelLister interface {
	// List lists all Channels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Channel, err error)
	// Channels returns an object that can list and get Channels.
	Channels(namespace string) ChannelNamespaceLister
	ChannelListerExpansion
}

// channelLister implements the ChannelLister interface.
type channelLister struct {
	indexer cache.Indexer
}

// NewChannelLister returns a new ChannelLister.
func NewChannelLister(indexer cache.Indexer) ChannelLister {
	return &channelLister{indexer: indexer}
}

// List lists all Channels in the indexer.
func (s *channelLister) List(selector labels.Selector) (ret []*v1alpha1.Channel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Channel))
	})
	return ret, err
}

// Channels returns an object that can list and get Channels.
func (s *channelLister) Channels(namespace string) ChannelNamespaceLister {
	return channelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ChannelNamespaceLister helps list and get Channels.
// All objects returned here must be treated as read-only.
type ChannelNamespaceLister interface {
	// List lists all Channels in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Channel, err error)
	// Get retrieves the Channel from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Channel, error)
	ChannelNamespaceListerExpansion
}

// channelNamespaceLister implements the ChannelNamespaceLister
// interface.
type channelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Channels in the indexer for a given namespace.
func (s channelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Channel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Channel))
	})
	return ret, err
}

// Get retrieves the Channel from the indexer for a given namespace and name.
func (s channelNamespaceLister) Get(name string) (*v1alpha1.Channel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("channel"), name)
	}
	return obj.(*v1alpha1.Channel), nil
}
