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

// PMDRoutesGetter has a method to return a PMDRouteInterface.
// A group's client should implement this interface.
type PMDRoutesGetter interface {
	PMDRoutes(namespace string) PMDRouteInterface
}

// PMDRouteInterface has methods to work with PMDRoute resources.
type PMDRouteInterface interface {
	Create(*v1.PMDRoute) (*v1.PMDRoute, error)
	Update(*v1.PMDRoute) (*v1.PMDRoute, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.PMDRoute, error)
	List(opts meta_v1.ListOptions) (*v1.PMDRouteList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PMDRoute, err error)
	PMDRouteExpansion
}

// pMDRoutes implements PMDRouteInterface
type pMDRoutes struct {
	client rest.Interface
	ns     string
}

// newPMDRoutes returns a PMDRoutes
func newPMDRoutes(c *PmdV1Client, namespace string) *pMDRoutes {
	return &pMDRoutes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pMDRoute, and returns the corresponding pMDRoute object, and an error if there is any.
func (c *pMDRoutes) Get(name string, options meta_v1.GetOptions) (result *v1.PMDRoute, err error) {
	result = &v1.PMDRoute{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pmdroutes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PMDRoutes that match those selectors.
func (c *pMDRoutes) List(opts meta_v1.ListOptions) (result *v1.PMDRouteList, err error) {
	result = &v1.PMDRouteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pmdroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pMDRoutes.
func (c *pMDRoutes) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pmdroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a pMDRoute and creates it.  Returns the server's representation of the pMDRoute, and an error, if there is any.
func (c *pMDRoutes) Create(pMDRoute *v1.PMDRoute) (result *v1.PMDRoute, err error) {
	result = &v1.PMDRoute{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pmdroutes").
		Body(pMDRoute).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pMDRoute and updates it. Returns the server's representation of the pMDRoute, and an error, if there is any.
func (c *pMDRoutes) Update(pMDRoute *v1.PMDRoute) (result *v1.PMDRoute, err error) {
	result = &v1.PMDRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pmdroutes").
		Name(pMDRoute.Name).
		Body(pMDRoute).
		Do().
		Into(result)
	return
}

// Delete takes name of the pMDRoute and deletes it. Returns an error if one occurs.
func (c *pMDRoutes) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pmdroutes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pMDRoutes) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pmdroutes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pMDRoute.
func (c *pMDRoutes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PMDRoute, err error) {
	result = &v1.PMDRoute{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pmdroutes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
