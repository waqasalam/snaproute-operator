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
package v1

import (
	v1 "snaproute-operator/pkg/apis/pmd/v1"
	scheme "snaproute-operator/pkg/client/clientset/versioned/scheme"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PMDAsNumbersGetter has a method to return a PMDAsNumberInterface.
// A group's client should implement this interface.
type PMDAsNumbersGetter interface {
	PMDAsNumbers(namespace string) PMDAsNumberInterface
}

// PMDAsNumberInterface has methods to work with PMDAsNumber resources.
type PMDAsNumberInterface interface {
	Create(*v1.PMDAsNumber) (*v1.PMDAsNumber, error)
	Update(*v1.PMDAsNumber) (*v1.PMDAsNumber, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.PMDAsNumber, error)
	List(opts meta_v1.ListOptions) (*v1.PMDAsNumberList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PMDAsNumber, err error)
	PMDAsNumberExpansion
}

// pMDAsNumbers implements PMDAsNumberInterface
type pMDAsNumbers struct {
	client rest.Interface
	ns     string
}

// newPMDAsNumbers returns a PMDAsNumbers
func newPMDAsNumbers(c *PmdV1Client, namespace string) *pMDAsNumbers {
	return &pMDAsNumbers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pMDAsNumber, and returns the corresponding pMDAsNumber object, and an error if there is any.
func (c *pMDAsNumbers) Get(name string, options meta_v1.GetOptions) (result *v1.PMDAsNumber, err error) {
	result = &v1.PMDAsNumber{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pmdasnumbers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PMDAsNumbers that match those selectors.
func (c *pMDAsNumbers) List(opts meta_v1.ListOptions) (result *v1.PMDAsNumberList, err error) {
	result = &v1.PMDAsNumberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pmdasnumbers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pMDAsNumbers.
func (c *pMDAsNumbers) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pmdasnumbers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a pMDAsNumber and creates it.  Returns the server's representation of the pMDAsNumber, and an error, if there is any.
func (c *pMDAsNumbers) Create(pMDAsNumber *v1.PMDAsNumber) (result *v1.PMDAsNumber, err error) {
	result = &v1.PMDAsNumber{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pmdasnumbers").
		Body(pMDAsNumber).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pMDAsNumber and updates it. Returns the server's representation of the pMDAsNumber, and an error, if there is any.
func (c *pMDAsNumbers) Update(pMDAsNumber *v1.PMDAsNumber) (result *v1.PMDAsNumber, err error) {
	result = &v1.PMDAsNumber{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pmdasnumbers").
		Name(pMDAsNumber.Name).
		Body(pMDAsNumber).
		Do().
		Into(result)
	return
}

// Delete takes name of the pMDAsNumber and deletes it. Returns an error if one occurs.
func (c *pMDAsNumbers) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pmdasnumbers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pMDAsNumbers) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pmdasnumbers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pMDAsNumber.
func (c *pMDAsNumbers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PMDAsNumber, err error) {
	result = &v1.PMDAsNumber{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pmdasnumbers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
