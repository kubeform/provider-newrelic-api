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

// FakeAlertConditions implements AlertConditionInterface
type FakeAlertConditions struct {
	Fake *FakeSyntheticsV1alpha1
	ns   string
}

var alertconditionsResource = schema.GroupVersionResource{Group: "synthetics.newrelic.kubeform.com", Version: "v1alpha1", Resource: "alertconditions"}

var alertconditionsKind = schema.GroupVersionKind{Group: "synthetics.newrelic.kubeform.com", Version: "v1alpha1", Kind: "AlertCondition"}

// Get takes name of the alertCondition, and returns the corresponding alertCondition object, and an error if there is any.
func (c *FakeAlertConditions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AlertCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alertconditionsResource, c.ns, name), &v1alpha1.AlertCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertCondition), err
}

// List takes label and field selectors, and returns the list of AlertConditions that match those selectors.
func (c *FakeAlertConditions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AlertConditionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alertconditionsResource, alertconditionsKind, c.ns, opts), &v1alpha1.AlertConditionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AlertConditionList{ListMeta: obj.(*v1alpha1.AlertConditionList).ListMeta}
	for _, item := range obj.(*v1alpha1.AlertConditionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alertConditions.
func (c *FakeAlertConditions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alertconditionsResource, c.ns, opts))

}

// Create takes the representation of a alertCondition and creates it.  Returns the server's representation of the alertCondition, and an error, if there is any.
func (c *FakeAlertConditions) Create(ctx context.Context, alertCondition *v1alpha1.AlertCondition, opts v1.CreateOptions) (result *v1alpha1.AlertCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alertconditionsResource, c.ns, alertCondition), &v1alpha1.AlertCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertCondition), err
}

// Update takes the representation of a alertCondition and updates it. Returns the server's representation of the alertCondition, and an error, if there is any.
func (c *FakeAlertConditions) Update(ctx context.Context, alertCondition *v1alpha1.AlertCondition, opts v1.UpdateOptions) (result *v1alpha1.AlertCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alertconditionsResource, c.ns, alertCondition), &v1alpha1.AlertCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertCondition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlertConditions) UpdateStatus(ctx context.Context, alertCondition *v1alpha1.AlertCondition, opts v1.UpdateOptions) (*v1alpha1.AlertCondition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alertconditionsResource, "status", c.ns, alertCondition), &v1alpha1.AlertCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertCondition), err
}

// Delete takes name of the alertCondition and deletes it. Returns an error if one occurs.
func (c *FakeAlertConditions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(alertconditionsResource, c.ns, name), &v1alpha1.AlertCondition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlertConditions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alertconditionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AlertConditionList{})
	return err
}

// Patch applies the patch and returns the patched alertCondition.
func (c *FakeAlertConditions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlertCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alertconditionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AlertCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertCondition), err
}