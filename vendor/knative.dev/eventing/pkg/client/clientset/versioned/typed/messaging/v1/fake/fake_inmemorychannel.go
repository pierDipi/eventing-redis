/*
Copyright 2021 The Knative Authors

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
	messagingv1 "knative.dev/eventing/pkg/apis/messaging/v1"
)

// FakeInMemoryChannels implements InMemoryChannelInterface
type FakeInMemoryChannels struct {
	Fake *FakeMessagingV1
	ns   string
}

var inmemorychannelsResource = schema.GroupVersionResource{Group: "messaging.knative.dev", Version: "v1", Resource: "inmemorychannels"}

var inmemorychannelsKind = schema.GroupVersionKind{Group: "messaging.knative.dev", Version: "v1", Kind: "InMemoryChannel"}

// Get takes name of the inMemoryChannel, and returns the corresponding inMemoryChannel object, and an error if there is any.
func (c *FakeInMemoryChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *messagingv1.InMemoryChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(inmemorychannelsResource, c.ns, name), &messagingv1.InMemoryChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*messagingv1.InMemoryChannel), err
}

// List takes label and field selectors, and returns the list of InMemoryChannels that match those selectors.
func (c *FakeInMemoryChannels) List(ctx context.Context, opts v1.ListOptions) (result *messagingv1.InMemoryChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(inmemorychannelsResource, inmemorychannelsKind, c.ns, opts), &messagingv1.InMemoryChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &messagingv1.InMemoryChannelList{ListMeta: obj.(*messagingv1.InMemoryChannelList).ListMeta}
	for _, item := range obj.(*messagingv1.InMemoryChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested inMemoryChannels.
func (c *FakeInMemoryChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(inmemorychannelsResource, c.ns, opts))

}

// Create takes the representation of a inMemoryChannel and creates it.  Returns the server's representation of the inMemoryChannel, and an error, if there is any.
func (c *FakeInMemoryChannels) Create(ctx context.Context, inMemoryChannel *messagingv1.InMemoryChannel, opts v1.CreateOptions) (result *messagingv1.InMemoryChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(inmemorychannelsResource, c.ns, inMemoryChannel), &messagingv1.InMemoryChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*messagingv1.InMemoryChannel), err
}

// Update takes the representation of a inMemoryChannel and updates it. Returns the server's representation of the inMemoryChannel, and an error, if there is any.
func (c *FakeInMemoryChannels) Update(ctx context.Context, inMemoryChannel *messagingv1.InMemoryChannel, opts v1.UpdateOptions) (result *messagingv1.InMemoryChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(inmemorychannelsResource, c.ns, inMemoryChannel), &messagingv1.InMemoryChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*messagingv1.InMemoryChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInMemoryChannels) UpdateStatus(ctx context.Context, inMemoryChannel *messagingv1.InMemoryChannel, opts v1.UpdateOptions) (*messagingv1.InMemoryChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(inmemorychannelsResource, "status", c.ns, inMemoryChannel), &messagingv1.InMemoryChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*messagingv1.InMemoryChannel), err
}

// Delete takes name of the inMemoryChannel and deletes it. Returns an error if one occurs.
func (c *FakeInMemoryChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(inmemorychannelsResource, c.ns, name, opts), &messagingv1.InMemoryChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInMemoryChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(inmemorychannelsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &messagingv1.InMemoryChannelList{})
	return err
}

// Patch applies the patch and returns the patched inMemoryChannel.
func (c *FakeInMemoryChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *messagingv1.InMemoryChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(inmemorychannelsResource, c.ns, name, pt, data, subresources...), &messagingv1.InMemoryChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*messagingv1.InMemoryChannel), err
}
