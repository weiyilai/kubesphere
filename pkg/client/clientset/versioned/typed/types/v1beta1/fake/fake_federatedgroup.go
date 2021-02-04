/*
Copyright 2020 The KubeSphere Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "kubesphere.io/kubesphere/pkg/apis/types/v1beta1"
)

// FakeFederatedGroups implements FederatedGroupInterface
type FakeFederatedGroups struct {
	Fake *FakeTypesV1beta1
}

var federatedgroupsResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatedgroups"}

var federatedgroupsKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedGroup"}

// Get takes name of the federatedGroup, and returns the corresponding federatedGroup object, and an error if there is any.
func (c *FakeFederatedGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FederatedGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(federatedgroupsResource, name), &v1beta1.FederatedGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedGroup), err
}

// List takes label and field selectors, and returns the list of FederatedGroups that match those selectors.
func (c *FakeFederatedGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FederatedGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(federatedgroupsResource, federatedgroupsKind, opts), &v1beta1.FederatedGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedGroupList{ListMeta: obj.(*v1beta1.FederatedGroupList).ListMeta}
	for _, item := range obj.(*v1beta1.FederatedGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedGroups.
func (c *FakeFederatedGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(federatedgroupsResource, opts))
}

// Create takes the representation of a federatedGroup and creates it.  Returns the server's representation of the federatedGroup, and an error, if there is any.
func (c *FakeFederatedGroups) Create(ctx context.Context, federatedGroup *v1beta1.FederatedGroup, opts v1.CreateOptions) (result *v1beta1.FederatedGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(federatedgroupsResource, federatedGroup), &v1beta1.FederatedGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedGroup), err
}

// Update takes the representation of a federatedGroup and updates it. Returns the server's representation of the federatedGroup, and an error, if there is any.
func (c *FakeFederatedGroups) Update(ctx context.Context, federatedGroup *v1beta1.FederatedGroup, opts v1.UpdateOptions) (result *v1beta1.FederatedGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(federatedgroupsResource, federatedGroup), &v1beta1.FederatedGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedGroups) UpdateStatus(ctx context.Context, federatedGroup *v1beta1.FederatedGroup, opts v1.UpdateOptions) (*v1beta1.FederatedGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(federatedgroupsResource, "status", federatedGroup), &v1beta1.FederatedGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedGroup), err
}

// Delete takes name of the federatedGroup and deletes it. Returns an error if one occurs.
func (c *FakeFederatedGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(federatedgroupsResource, name), &v1beta1.FederatedGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(federatedgroupsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedGroupList{})
	return err
}

// Patch applies the patch and returns the patched federatedGroup.
func (c *FakeFederatedGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(federatedgroupsResource, name, pt, data, subresources...), &v1beta1.FederatedGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedGroup), err
}