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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-newrelic-api/apis/nrql/v1alpha1"
	scheme "kubeform.dev/provider-newrelic-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DropRulesGetter has a method to return a DropRuleInterface.
// A group's client should implement this interface.
type DropRulesGetter interface {
	DropRules(namespace string) DropRuleInterface
}

// DropRuleInterface has methods to work with DropRule resources.
type DropRuleInterface interface {
	Create(ctx context.Context, dropRule *v1alpha1.DropRule, opts v1.CreateOptions) (*v1alpha1.DropRule, error)
	Update(ctx context.Context, dropRule *v1alpha1.DropRule, opts v1.UpdateOptions) (*v1alpha1.DropRule, error)
	UpdateStatus(ctx context.Context, dropRule *v1alpha1.DropRule, opts v1.UpdateOptions) (*v1alpha1.DropRule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DropRule, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DropRuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DropRule, err error)
	DropRuleExpansion
}

// dropRules implements DropRuleInterface
type dropRules struct {
	client rest.Interface
	ns     string
}

// newDropRules returns a DropRules
func newDropRules(c *NrqlV1alpha1Client, namespace string) *dropRules {
	return &dropRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dropRule, and returns the corresponding dropRule object, and an error if there is any.
func (c *dropRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DropRule, err error) {
	result = &v1alpha1.DropRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("droprules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DropRules that match those selectors.
func (c *dropRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DropRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DropRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("droprules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dropRules.
func (c *dropRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("droprules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dropRule and creates it.  Returns the server's representation of the dropRule, and an error, if there is any.
func (c *dropRules) Create(ctx context.Context, dropRule *v1alpha1.DropRule, opts v1.CreateOptions) (result *v1alpha1.DropRule, err error) {
	result = &v1alpha1.DropRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("droprules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dropRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dropRule and updates it. Returns the server's representation of the dropRule, and an error, if there is any.
func (c *dropRules) Update(ctx context.Context, dropRule *v1alpha1.DropRule, opts v1.UpdateOptions) (result *v1alpha1.DropRule, err error) {
	result = &v1alpha1.DropRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("droprules").
		Name(dropRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dropRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dropRules) UpdateStatus(ctx context.Context, dropRule *v1alpha1.DropRule, opts v1.UpdateOptions) (result *v1alpha1.DropRule, err error) {
	result = &v1alpha1.DropRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("droprules").
		Name(dropRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dropRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dropRule and deletes it. Returns an error if one occurs.
func (c *dropRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("droprules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dropRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("droprules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dropRule.
func (c *dropRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DropRule, err error) {
	result = &v1alpha1.DropRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("droprules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}