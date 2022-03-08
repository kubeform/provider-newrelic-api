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
	v1alpha1 "kubeform.dev/provider-newrelic-api/apis/service/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LevelLister helps list Levels.
// All objects returned here must be treated as read-only.
type LevelLister interface {
	// List lists all Levels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Level, err error)
	// Levels returns an object that can list and get Levels.
	Levels(namespace string) LevelNamespaceLister
	LevelListerExpansion
}

// levelLister implements the LevelLister interface.
type levelLister struct {
	indexer cache.Indexer
}

// NewLevelLister returns a new LevelLister.
func NewLevelLister(indexer cache.Indexer) LevelLister {
	return &levelLister{indexer: indexer}
}

// List lists all Levels in the indexer.
func (s *levelLister) List(selector labels.Selector) (ret []*v1alpha1.Level, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Level))
	})
	return ret, err
}

// Levels returns an object that can list and get Levels.
func (s *levelLister) Levels(namespace string) LevelNamespaceLister {
	return levelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LevelNamespaceLister helps list and get Levels.
// All objects returned here must be treated as read-only.
type LevelNamespaceLister interface {
	// List lists all Levels in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Level, err error)
	// Get retrieves the Level from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Level, error)
	LevelNamespaceListerExpansion
}

// levelNamespaceLister implements the LevelNamespaceLister
// interface.
type levelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Levels in the indexer for a given namespace.
func (s levelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Level, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Level))
	})
	return ret, err
}

// Get retrieves the Level from the indexer for a given namespace and name.
func (s levelNamespaceLister) Get(name string) (*v1alpha1.Level, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("level"), name)
	}
	return obj.(*v1alpha1.Level), nil
}