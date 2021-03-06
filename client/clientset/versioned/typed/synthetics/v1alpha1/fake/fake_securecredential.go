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

	v1alpha1 "kubeform.dev/provider-newrelic-api/apis/synthetics/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSecureCredentials implements SecureCredentialInterface
type FakeSecureCredentials struct {
	Fake *FakeSyntheticsV1alpha1
	ns   string
}

var securecredentialsResource = schema.GroupVersionResource{Group: "synthetics.newrelic.kubeform.com", Version: "v1alpha1", Resource: "securecredentials"}

var securecredentialsKind = schema.GroupVersionKind{Group: "synthetics.newrelic.kubeform.com", Version: "v1alpha1", Kind: "SecureCredential"}

// Get takes name of the secureCredential, and returns the corresponding secureCredential object, and an error if there is any.
func (c *FakeSecureCredentials) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SecureCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(securecredentialsResource, c.ns, name), &v1alpha1.SecureCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecureCredential), err
}

// List takes label and field selectors, and returns the list of SecureCredentials that match those selectors.
func (c *FakeSecureCredentials) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SecureCredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(securecredentialsResource, securecredentialsKind, c.ns, opts), &v1alpha1.SecureCredentialList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecureCredentialList{ListMeta: obj.(*v1alpha1.SecureCredentialList).ListMeta}
	for _, item := range obj.(*v1alpha1.SecureCredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secureCredentials.
func (c *FakeSecureCredentials) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(securecredentialsResource, c.ns, opts))

}

// Create takes the representation of a secureCredential and creates it.  Returns the server's representation of the secureCredential, and an error, if there is any.
func (c *FakeSecureCredentials) Create(ctx context.Context, secureCredential *v1alpha1.SecureCredential, opts v1.CreateOptions) (result *v1alpha1.SecureCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(securecredentialsResource, c.ns, secureCredential), &v1alpha1.SecureCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecureCredential), err
}

// Update takes the representation of a secureCredential and updates it. Returns the server's representation of the secureCredential, and an error, if there is any.
func (c *FakeSecureCredentials) Update(ctx context.Context, secureCredential *v1alpha1.SecureCredential, opts v1.UpdateOptions) (result *v1alpha1.SecureCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(securecredentialsResource, c.ns, secureCredential), &v1alpha1.SecureCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecureCredential), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecureCredentials) UpdateStatus(ctx context.Context, secureCredential *v1alpha1.SecureCredential, opts v1.UpdateOptions) (*v1alpha1.SecureCredential, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(securecredentialsResource, "status", c.ns, secureCredential), &v1alpha1.SecureCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecureCredential), err
}

// Delete takes name of the secureCredential and deletes it. Returns an error if one occurs.
func (c *FakeSecureCredentials) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(securecredentialsResource, c.ns, name), &v1alpha1.SecureCredential{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecureCredentials) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(securecredentialsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecureCredentialList{})
	return err
}

// Patch applies the patch and returns the patched secureCredential.
func (c *FakeSecureCredentials) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecureCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(securecredentialsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SecureCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecureCredential), err
}
