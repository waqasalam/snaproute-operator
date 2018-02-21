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

// FakePMDAsNumbers implements PMDAsNumberInterface
type FakePMDAsNumbers struct {
	Fake *FakePmdV1
	ns   string
}

var pmdasnumbersResource = schema.GroupVersionResource{Group: "pmd.snaproute.com", Version: "v1", Resource: "pmdasnumbers"}

var pmdasnumbersKind = schema.GroupVersionKind{Group: "pmd.snaproute.com", Version: "v1", Kind: "PMDAsNumber"}

// Get takes name of the pMDAsNumber, and returns the corresponding pMDAsNumber object, and an error if there is any.
func (c *FakePMDAsNumbers) Get(name string, options v1.GetOptions) (result *pmd_v1.PMDAsNumber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pmdasnumbersResource, c.ns, name), &pmd_v1.PMDAsNumber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pmd_v1.PMDAsNumber), err
}

// List takes label and field selectors, and returns the list of PMDAsNumbers that match those selectors.
func (c *FakePMDAsNumbers) List(opts v1.ListOptions) (result *pmd_v1.PMDAsNumberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pmdasnumbersResource, pmdasnumbersKind, c.ns, opts), &pmd_v1.PMDAsNumberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &pmd_v1.PMDAsNumberList{}
	for _, item := range obj.(*pmd_v1.PMDAsNumberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pMDAsNumbers.
func (c *FakePMDAsNumbers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pmdasnumbersResource, c.ns, opts))

}

// Create takes the representation of a pMDAsNumber and creates it.  Returns the server's representation of the pMDAsNumber, and an error, if there is any.
func (c *FakePMDAsNumbers) Create(pMDAsNumber *pmd_v1.PMDAsNumber) (result *pmd_v1.PMDAsNumber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pmdasnumbersResource, c.ns, pMDAsNumber), &pmd_v1.PMDAsNumber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pmd_v1.PMDAsNumber), err
}

// Update takes the representation of a pMDAsNumber and updates it. Returns the server's representation of the pMDAsNumber, and an error, if there is any.
func (c *FakePMDAsNumbers) Update(pMDAsNumber *pmd_v1.PMDAsNumber) (result *pmd_v1.PMDAsNumber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pmdasnumbersResource, c.ns, pMDAsNumber), &pmd_v1.PMDAsNumber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pmd_v1.PMDAsNumber), err
}

// Delete takes name of the pMDAsNumber and deletes it. Returns an error if one occurs.
func (c *FakePMDAsNumbers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pmdasnumbersResource, c.ns, name), &pmd_v1.PMDAsNumber{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePMDAsNumbers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pmdasnumbersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &pmd_v1.PMDAsNumberList{})
	return err
}

// Patch applies the patch and returns the patched pMDAsNumber.
func (c *FakePMDAsNumbers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *pmd_v1.PMDAsNumber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pmdasnumbersResource, c.ns, name, data, subresources...), &pmd_v1.PMDAsNumber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pmd_v1.PMDAsNumber), err
}
