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

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/provider-newrelic-api/apis/entity/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTagses implements TagsInterface
type FakeTagses struct {
	Fake *FakeEntityV1alpha1
	ns   string
}

var tagsesResource = schema.GroupVersionResource{Group: "entity.newrelic.kubeform.com", Version: "v1alpha1", Resource: "tagses"}

var tagsesKind = schema.GroupVersionKind{Group: "entity.newrelic.kubeform.com", Version: "v1alpha1", Kind: "Tags"}

// Get takes name of the tags, and returns the corresponding tags object, and an error if there is any.
func (c *FakeTagses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Tags, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tagsesResource, c.ns, name), &v1alpha1.Tags{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tags), err
}

// List takes label and field selectors, and returns the list of Tagses that match those selectors.
func (c *FakeTagses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TagsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tagsesResource, tagsesKind, c.ns, opts), &v1alpha1.TagsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TagsList{ListMeta: obj.(*v1alpha1.TagsList).ListMeta}
	for _, item := range obj.(*v1alpha1.TagsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tagses.
func (c *FakeTagses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tagsesResource, c.ns, opts))

}

// Create takes the representation of a tags and creates it.  Returns the server's representation of the tags, and an error, if there is any.
func (c *FakeTagses) Create(ctx context.Context, tags *v1alpha1.Tags, opts v1.CreateOptions) (result *v1alpha1.Tags, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tagsesResource, c.ns, tags), &v1alpha1.Tags{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tags), err
}

// Update takes the representation of a tags and updates it. Returns the server's representation of the tags, and an error, if there is any.
func (c *FakeTagses) Update(ctx context.Context, tags *v1alpha1.Tags, opts v1.UpdateOptions) (result *v1alpha1.Tags, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tagsesResource, c.ns, tags), &v1alpha1.Tags{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tags), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTagses) UpdateStatus(ctx context.Context, tags *v1alpha1.Tags, opts v1.UpdateOptions) (*v1alpha1.Tags, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tagsesResource, "status", c.ns, tags), &v1alpha1.Tags{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tags), err
}

// Delete takes name of the tags and deletes it. Returns an error if one occurs.
func (c *FakeTagses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tagsesResource, c.ns, name), &v1alpha1.Tags{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTagses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tagsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TagsList{})
	return err
}

// Patch applies the patch and returns the patched tags.
func (c *FakeTagses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Tags, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tagsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Tags{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tags), err
}