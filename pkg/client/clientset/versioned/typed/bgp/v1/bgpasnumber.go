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
	v1 "snaproute-operator/pkg/apis/bgp/v1"
	scheme "snaproute-operator/pkg/client/clientset/versioned/scheme"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BGPAsNumbersGetter has a method to return a BGPAsNumberInterface.
// A group's client should implement this interface.
type BGPAsNumbersGetter interface {
	BGPAsNumbers(namespace string) BGPAsNumberInterface
}

// BGPAsNumberInterface has methods to work with BGPAsNumber resources.
type BGPAsNumberInterface interface {
	Create(*v1.BGPAsNumber) (*v1.BGPAsNumber, error)
	Update(*v1.BGPAsNumber) (*v1.BGPAsNumber, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.BGPAsNumber, error)
	List(opts meta_v1.ListOptions) (*v1.BGPAsNumberList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.BGPAsNumber, err error)
	BGPAsNumberExpansion
}

// bGPAsNumbers implements BGPAsNumberInterface
type bGPAsNumbers struct {
	client rest.Interface
	ns     string
}

// newBGPAsNumbers returns a BGPAsNumbers
func newBGPAsNumbers(c *BgpV1Client, namespace string) *bGPAsNumbers {
	return &bGPAsNumbers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bGPAsNumber, and returns the corresponding bGPAsNumber object, and an error if there is any.
func (c *bGPAsNumbers) Get(name string, options meta_v1.GetOptions) (result *v1.BGPAsNumber, err error) {
	result = &v1.BGPAsNumber{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bgpasnumbers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BGPAsNumbers that match those selectors.
func (c *bGPAsNumbers) List(opts meta_v1.ListOptions) (result *v1.BGPAsNumberList, err error) {
	result = &v1.BGPAsNumberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bgpasnumbers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bGPAsNumbers.
func (c *bGPAsNumbers) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bgpasnumbers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a bGPAsNumber and creates it.  Returns the server's representation of the bGPAsNumber, and an error, if there is any.
func (c *bGPAsNumbers) Create(bGPAsNumber *v1.BGPAsNumber) (result *v1.BGPAsNumber, err error) {
	result = &v1.BGPAsNumber{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bgpasnumbers").
		Body(bGPAsNumber).
		Do().
		Into(result)
	return
}

// Update takes the representation of a bGPAsNumber and updates it. Returns the server's representation of the bGPAsNumber, and an error, if there is any.
func (c *bGPAsNumbers) Update(bGPAsNumber *v1.BGPAsNumber) (result *v1.BGPAsNumber, err error) {
	result = &v1.BGPAsNumber{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bgpasnumbers").
		Name(bGPAsNumber.Name).
		Body(bGPAsNumber).
		Do().
		Into(result)
	return
}

// Delete takes name of the bGPAsNumber and deletes it. Returns an error if one occurs.
func (c *bGPAsNumbers) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bgpasnumbers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bGPAsNumbers) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bgpasnumbers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched bGPAsNumber.
func (c *bGPAsNumbers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.BGPAsNumber, err error) {
	result = &v1.BGPAsNumber{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bgpasnumbers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
