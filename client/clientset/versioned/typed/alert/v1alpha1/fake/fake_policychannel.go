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

	v1alpha1 "kubeform.dev/provider-newrelic-api/apis/alert/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePolicyChannels implements PolicyChannelInterface
type FakePolicyChannels struct {
	Fake *FakeAlertV1alpha1
	ns   string
}

var policychannelsResource = schema.GroupVersionResource{Group: "alert.newrelic.kubeform.com", Version: "v1alpha1", Resource: "policychannels"}

var policychannelsKind = schema.GroupVersionKind{Group: "alert.newrelic.kubeform.com", Version: "v1alpha1", Kind: "PolicyChannel"}

// Get takes name of the policyChannel, and returns the corresponding policyChannel object, and an error if there is any.
func (c *FakePolicyChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PolicyChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(policychannelsResource, c.ns, name), &v1alpha1.PolicyChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyChannel), err
}

// List takes label and field selectors, and returns the list of PolicyChannels that match those selectors.
func (c *FakePolicyChannels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PolicyChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(policychannelsResource, policychannelsKind, c.ns, opts), &v1alpha1.PolicyChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PolicyChannelList{ListMeta: obj.(*v1alpha1.PolicyChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.PolicyChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested policyChannels.
func (c *FakePolicyChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(policychannelsResource, c.ns, opts))

}

// Create takes the representation of a policyChannel and creates it.  Returns the server's representation of the policyChannel, and an error, if there is any.
func (c *FakePolicyChannels) Create(ctx context.Context, policyChannel *v1alpha1.PolicyChannel, opts v1.CreateOptions) (result *v1alpha1.PolicyChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(policychannelsResource, c.ns, policyChannel), &v1alpha1.PolicyChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyChannel), err
}

// Update takes the representation of a policyChannel and updates it. Returns the server's representation of the policyChannel, and an error, if there is any.
func (c *FakePolicyChannels) Update(ctx context.Context, policyChannel *v1alpha1.PolicyChannel, opts v1.UpdateOptions) (result *v1alpha1.PolicyChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(policychannelsResource, c.ns, policyChannel), &v1alpha1.PolicyChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePolicyChannels) UpdateStatus(ctx context.Context, policyChannel *v1alpha1.PolicyChannel, opts v1.UpdateOptions) (*v1alpha1.PolicyChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(policychannelsResource, "status", c.ns, policyChannel), &v1alpha1.PolicyChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyChannel), err
}

// Delete takes name of the policyChannel and deletes it. Returns an error if one occurs.
func (c *FakePolicyChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(policychannelsResource, c.ns, name), &v1alpha1.PolicyChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePolicyChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(policychannelsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PolicyChannelList{})
	return err
}

// Patch applies the patch and returns the patched policyChannel.
func (c *FakePolicyChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PolicyChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(policychannelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PolicyChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyChannel), err
}
