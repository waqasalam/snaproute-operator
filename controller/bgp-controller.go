package main

import (
	"bgp-crd/client"
	"bgp-crd/crd"
	"flag"
	"fmt"
	apiextcs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

// return rest config, if path not specified assume in cluster config
func GetClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}

func main() {

	//	kubeconf := flag.String("kubeconf", "admin.conf", "Path to a kube config. Only required if out-of-cluster.")
	flag.Parse()
	kubeconf := ""
	config, err := GetClientConfig(kubeconf)
	if err != nil {
		panic(err.Error())
	}

	// create clientset and create our CRD, this only need to run once
	clientset, err := apiextcs.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// note: if the CRD exist our CreateCRD function is set to exit without an error
	err = crd.CreateCRD(clientset)
	if err != nil {
		panic(err)
	}

	// Wait for the CRD to be created before we use it (only needed if its a new one)
	time.Sleep(3 * time.Second)

	// Create a new clientset which include our CRD schema
	crdcs, scheme, err := crd.NewClient(config)
	if err != nil {
		panic(err)
	}

	// Create a CRD client interface
	crdclient := client.CrdClient(crdcs, scheme, "default")

	// Create a new BGPAsNumber object and write to k8s you can use
	// bgp.yaml
	bgpasnumber := &crd.BGPAsNumber{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:   "bgpasnumber1",
			Labels: map[string]string{"mylabel": "test"},
		},
		Spec: crd.BGPAsNumberSpec{
			AsNumber: "one",
			Enable:   true,
		},
		Status: crd.BGPAsNumberStatus{
			State:   "created",
			Message: "Created, not processed yet",
		},
	}

	result, err := crdclient.Create(bgpasnumber)
	if err == nil {
		fmt.Printf("CREATED: %#v\n", result)
	} else if apierrors.IsAlreadyExists(err) {
		fmt.Printf("ALREADY EXISTS: %#v\n", result)
	} else {
		panic(err)
	}

	// List all BGP AsNumber objects
	items, err := crdclient.List(meta_v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("List:\n%s\n", items)

	// BGP Controller
	// Watch for changes in BGP objects and fire Add, Delete, Update callbacks
	_, controller := cache.NewInformer(
		crdclient.NewListWatch(),
		&crd.BGPAsNumber{},
		time.Minute*10,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Printf("add: %s \n", obj)
				ex := obj.(*crd.BGPAsNumber)
				fmt.Println("AsNumber", ex.Spec.AsNumber, "Enable", ex.Spec.Enable)
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Printf("delete: %s \n", obj)
				ex := obj.(*crd.BGPAsNumber)
				fmt.Printf("AsNumber", ex.Spec.AsNumber, "Enable", ex.Spec.Enable)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Printf("Update old: %s \n      New: %s\n", oldObj, newObj)
				ex := newObj.(*crd.BGPAsNumber)
				fmt.Printf("AsNumber", ex.Spec.AsNumber, "Enable", ex.Spec.Enable)
			},
		},
	)

	stop := make(chan struct{})
	go controller.Run(stop)

	// Wait forever
	select {}
}
