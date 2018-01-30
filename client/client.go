package client

import (
	"snaproute-operator/crd"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

// This file implement all the (CRUD) client methods we need to access our CRD object

func CrdClient(cl *rest.RESTClient, scheme *runtime.Scheme, namespace string) *crdclient {
	return &crdclient{cl: cl, ns: namespace, plural: crd.CRDPlural,
		codec: runtime.NewParameterCodec(scheme)}
}

type crdclient struct {
	cl     *rest.RESTClient
	ns     string
	plural string
	codec  runtime.ParameterCodec
}

// implement
// Create takes the representation of BGPAsNumber and creates it.  Returns the server's representation of the podDisruptionBudget, and an error, if there is any.
func (f *crdclient) Create(obj *crd.BGPAsNumber) (*crd.BGPAsNumber, error) {
	var result crd.BGPAsNumber
	err := f.cl.Post().
		Namespace(f.ns).Resource(f.plural).
		Body(obj).Do().Into(&result)
	return &result, err
}

// Update takes the representation of BGPAsNumber and updates it. Returns the server's representation of the podDisruptionBudget, and an error, if there is any.
func (f *crdclient) Update(obj *crd.BGPAsNumber) (*crd.BGPAsNumber, error) {
	var result crd.BGPAsNumber
	err := f.cl.Put().
		Namespace(f.ns).Resource(f.plural).
		Body(obj).Do().Into(&result)
	return &result, err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
/*
func (f *crdclient) UpdateStatus(obj *crd.BGPAsNumber) (*crd.BGPAsNumber, error) {
	var result crd.BGPAsNumber
	err = f.cl.Put().
		Namespace(f.ns).
		Resource(f.plural).
		Name(obj.Name).
		SubResource("status").
		Body(obj).
		Do().
		Into(&result)
	return
}
*/
// Delete takes name of the BGPAsNumber and deletes it. Returns an error if one occurs.
func (f *crdclient) Delete(name string, options *meta_v1.DeleteOptions) error {
	return f.cl.Delete().
		Namespace(f.ns).Resource(f.plural).
		Name(name).Body(options).Do().
		Error()
}

func (f *crdclient) Get(name string) (*crd.BGPAsNumber, error) {
	var result crd.BGPAsNumber
	err := f.cl.Get().
		Namespace(f.ns).Resource(f.plural).
		Name(name).Do().Into(&result)
	return &result, err
}

func (f *crdclient) List(opts meta_v1.ListOptions) (*crd.BGPAsNumberList, error) {
	var result crd.BGPAsNumberList
	err := f.cl.Get().
		Namespace(f.ns).Resource(f.plural).
		VersionedParams(&opts, f.codec).
		Do().Into(&result)
	return &result, err
}

// implement
func (f *crdclient) NewListWatch() *cache.ListWatch {
	return cache.NewListWatchFromClient(f.cl, f.plural, f.ns, fields.Everything())
}
