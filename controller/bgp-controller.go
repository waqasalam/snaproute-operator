package main

import (
	"flag"
	"fmt"
	apiextcs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"snaproute-operator/client"
	"snaproute-operator/crd"
)

// return rest config, if path not specified assume in cluster config
func GetClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}

func main() {
	fmt.Println("Start the CRD")
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
	err = crd.CreateCRD(clientset, "bgpasnumber")
	if err != nil {
		panic(err)
	}

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
			Message: "Created, not -- processed yet",
		},
	}

	result, err := crdclient.Create(bgpasnumber)
	if err == nil {
		fmt.Printf("CREATED: %#v\n", result)
	} else if apierrors.IsAlreadyExists(err) {
		fmt.Printf("ALREADY-- EXISTS: %#v\n", result)
	} else {
		panic(err)
	}

	// List all BGP AsNumber objects
	items, err := crdclient.List(meta_v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Do list operations :List:\n%s\n", items)
	fmt.Println("List operation done")
	// BGP Controller
	// Watch for changes in BGP objects and fire Add, Delete, Update callbacks

	controller := cache.NewSharedIndexInformer(
		crdclient.NewListWatch(),
		&crd.BGPAsNumber{},
		0,
		cache.Indexers{},
	)
	controller.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("add -- crd shared: %s \n", obj)
			ex := obj.(*crd.BGPAsNumber)
			fmt.Println("AsNumber", ex.Spec.AsNumber, "Enable", ex.Spec.Enable)
			newex := ex.DeepCopy()
			newex.Status.Message = "Processed in handler"
			newex.Status.State = "Created"
			fmt.Printf("newex %+v\n", newex)
			newex.Spec.AsNumber = "555"
			if _, e := crdclient.Update(newex); e != nil {
				fmt.Println("update satus error", e)
			}
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Printf("delete: %s shared \n", obj)
			ex := obj.(*crd.BGPAsNumber)
			fmt.Printf("AsNumber", ex.Spec.AsNumber, "Enable", ex.Spec.Enable)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Printf("Update -- old: %s \n      New: %s\n", oldObj, newObj)
			ex := newObj.(*crd.BGPAsNumber)
			fmt.Printf("AsNumber", ex.Spec.AsNumber, "Enable", ex.Spec.Enable)
			newex := ex.DeepCopy()
			newex.Status.Message = "Processed in handler"
			newex.Status.State = "Created"
			newex.Spec.AsNumber = "555"
			fmt.Printf("newex %+v\n", newex)
			if _, e := crdclient.Update(newex); e != nil {
				fmt.Println("update satus error", e)
			}
		},
	},
	)
	stop := make(chan struct{})
	go controller.Run(stop)

	// Wait forever
	select {}
}
