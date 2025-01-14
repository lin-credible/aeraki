// Copyright Aeraki Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	v1alpha1 "github.com/aeraki-mesh/aeraki/client-go/pkg/apis/metaprotocol/v1alpha1"
)

// FakeMetaRouters implements MetaRouterInterface
type FakeMetaRouters struct {
	Fake *FakeMetaprotocolV1alpha1
	ns   string
}

var metaroutersResource = schema.GroupVersionResource{Group: "metaprotocol.aeraki.io", Version: "v1alpha1", Resource: "metarouters"}

var metaroutersKind = schema.GroupVersionKind{Group: "metaprotocol.aeraki.io", Version: "v1alpha1", Kind: "MetaRouter"}

// Get takes name of the metaRouter, and returns the corresponding metaRouter object, and an error if there is any.
func (c *FakeMetaRouters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MetaRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(metaroutersResource, c.ns, name), &v1alpha1.MetaRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetaRouter), err
}

// List takes label and field selectors, and returns the list of MetaRouters that match those selectors.
func (c *FakeMetaRouters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MetaRouterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(metaroutersResource, metaroutersKind, c.ns, opts), &v1alpha1.MetaRouterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MetaRouterList{ListMeta: obj.(*v1alpha1.MetaRouterList).ListMeta}
	for _, item := range obj.(*v1alpha1.MetaRouterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested metaRouters.
func (c *FakeMetaRouters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(metaroutersResource, c.ns, opts))

}

// Create takes the representation of a metaRouter and creates it.  Returns the server's representation of the metaRouter, and an error, if there is any.
func (c *FakeMetaRouters) Create(ctx context.Context, metaRouter *v1alpha1.MetaRouter, opts v1.CreateOptions) (result *v1alpha1.MetaRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(metaroutersResource, c.ns, metaRouter), &v1alpha1.MetaRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetaRouter), err
}

// Update takes the representation of a metaRouter and updates it. Returns the server's representation of the metaRouter, and an error, if there is any.
func (c *FakeMetaRouters) Update(ctx context.Context, metaRouter *v1alpha1.MetaRouter, opts v1.UpdateOptions) (result *v1alpha1.MetaRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(metaroutersResource, c.ns, metaRouter), &v1alpha1.MetaRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetaRouter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMetaRouters) UpdateStatus(ctx context.Context, metaRouter *v1alpha1.MetaRouter, opts v1.UpdateOptions) (*v1alpha1.MetaRouter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(metaroutersResource, "status", c.ns, metaRouter), &v1alpha1.MetaRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetaRouter), err
}

// Delete takes name of the metaRouter and deletes it. Returns an error if one occurs.
func (c *FakeMetaRouters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(metaroutersResource, c.ns, name), &v1alpha1.MetaRouter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMetaRouters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(metaroutersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MetaRouterList{})
	return err
}

// Patch applies the patch and returns the patched metaRouter.
func (c *FakeMetaRouters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MetaRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(metaroutersResource, c.ns, name, pt, data, subresources...), &v1alpha1.MetaRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetaRouter), err
}
