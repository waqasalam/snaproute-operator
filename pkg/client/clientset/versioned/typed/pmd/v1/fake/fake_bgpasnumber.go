/*
Copyright 2018 The Voyager Authors.

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
package fake

import (
	pmd_v1 "snaproute-operator/pkg/apis/pmd/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBGPAsNumbers implements BGPAsNumberInterface
type FakeBGPAsNumbers struct {
	Fake *FakePmdV1
	ns   string
}

var bgpasnumbersResource = schema.GroupVersionResource{Group: "pmd", Version: "v1", Resource: "bgpasnumbers"}

var bgpasnumbersKind = schema.GroupVersionKind{Group: "pmd", Version: "v1", Kind: "BGPAsNumber"}

// Get takes name of the bGPAsNumber, and returns the corresponding bGPAsNumber object, and an error if there is any.
func (c *FakeBGPAsNumbers) Get(name string, options v1.GetOptions) (result *pmd_v1.BGPAsNumber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bgpasnumbersResource, c.ns, name), &pmd_v1.BGPAsNumber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pmd_v1.BGPAsNumber), err
}

// List takes label and field selectors, and returns the list of BGPAsNumbers that match those selectors.
func (c *FakeBGPAsNumbers) List(opts v1.ListOptions) (result *pmd_v1.BGPAsNumberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bgpasnumbersResource, bgpasnumbersKind, c.ns, opts), &pmd_v1.BGPAsNumberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &pmd_v1.BGPAsNumberList{}
	for _, item := range obj.(*pmd_v1.BGPAsNumberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bGPAsNumbers.
func (c *FakeBGPAsNumbers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bgpasnumbersResource, c.ns, opts))

}

// Create takes the representation of a bGPAsNumber and creates it.  Returns the server's representation of the bGPAsNumber, and an error, if there is any.
func (c *FakeBGPAsNumbers) Create(bGPAsNumber *pmd_v1.BGPAsNumber) (result *pmd_v1.BGPAsNumber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bgpasnumbersResource, c.ns, bGPAsNumber), &pmd_v1.BGPAsNumber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pmd_v1.BGPAsNumber), err
}

// Update takes the representation of a bGPAsNumber and updates it. Returns the server's representation of the bGPAsNumber, and an error, if there is any.
func (c *FakeBGPAsNumbers) Update(bGPAsNumber *pmd_v1.BGPAsNumber) (result *pmd_v1.BGPAsNumber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bgpasnumbersResource, c.ns, bGPAsNumber), &pmd_v1.BGPAsNumber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pmd_v1.BGPAsNumber), err
}

// Delete takes name of the bGPAsNumber and deletes it. Returns an error if one occurs.
func (c *FakeBGPAsNumbers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bgpasnumbersResource, c.ns, name), &pmd_v1.BGPAsNumber{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBGPAsNumbers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bgpasnumbersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &pmd_v1.BGPAsNumberList{})
	return err
}

// Patch applies the patch and returns the patched bGPAsNumber.
func (c *FakeBGPAsNumbers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *pmd_v1.BGPAsNumber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bgpasnumbersResource, c.ns, name, data, subresources...), &pmd_v1.BGPAsNumber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pmd_v1.BGPAsNumber), err
}
