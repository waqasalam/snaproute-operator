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
	"snaproute-operator/pkg/apis/bgp/v1"
	clientset "snaproute-operator/pkg/client/clientset/versioned"
	scheme "snaproute-operator/pkg/client/clientset/versioned/scheme"
	informers "snaproute-operator/pkg/client/informers/externalversions"
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
	fmt.Println("Start the CRD")
	//	kubeconf := flag.String("kubeconf", "admin.conf", "Path to a kube config. Only required if out-of-cluster.")
	flag.Parse()
	kubeconf := ""
	config, err := GetClientConfig(kubeconf)
	if err != nil {
		panic(err.Error())
	}

	// create clientset and create our CRD, this only need to run once
	clientset, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// create clientset and create our CRD, this only need to run once
	clt, err := apiextcs.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// note: if the CRD exist our CreateCRD function is set to exit without an error
	err = v1.CreateCRD(clt)
	if err != nil {
		panic(err)
	}

	scheme.AddToScheme(scheme.Scheme)
	bgpInformerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	informer := bgpInformerFactory.Bgp().V1().BGPAsNumbers()
	// Create a new BGPAsNumber object and write to k8s you can use
	// bgp.yaml
	bgpasnumber := &v1.BGPAsNumber{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:   "bgpasnumber1",
			Labels: map[string]string{"mylabel": "test"},
		},
		Spec: v1.BGPAsNumberSpec{
			AsNumber: "one",
			Enable:   true,
		},
		Status: v1.BGPAsNumberStatus{
			State:   "created",
			Message: "Created, not -- processed yet",
		},
	}
	fmt.Printf("crd create")
	result, err := clientset.BgpV1().BGPAsNumbers("default").Create(bgpasnumber)
	if err == nil {
		fmt.Printf("CREATED: %#v\n", result)
	} else if apierrors.IsAlreadyExists(err) {
		fmt.Printf("ALREADY-- EXISTS: %#v\n", result)
	} else {
		panic(err)
	}
	fmt.Printf("crd creation done")

	// List all BGP AsNumber objects
	items, err := clientset.BgpV1().BGPAsNumbers("default").List(meta_v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Do list operations :List:\n%s\n", items)
	fmt.Println("List operation done")
	// BGP Controller
	// Watch for changes in BGP objects and fire Add, Delete, Update callbacks

	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("add -- crd shared: %s \n", obj)
			ex := obj.(*v1.BGPAsNumber)
			fmt.Println("AsNumber", ex.Spec.AsNumber, "Enable", ex.Spec.Enable)
			newex := ex.DeepCopy()
			newex.Status.Message = "Processed in handler"
			newex.Status.State = "Created"
			fmt.Printf("newex %+v\n", newex)
			newex.Spec.AsNumber = "555"
			if _, e := clientset.BgpV1().BGPAsNumbers("default").Update(newex); e != nil {
				fmt.Println("update satus error", e)
			}
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Printf("delete: %s shared \n", obj)
			ex := obj.(*v1.BGPAsNumber)
			fmt.Printf("AsNumber", ex.Spec.AsNumber, "Enable", ex.Spec.Enable)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Printf("Update -- old: %s \n      New: %s\n", oldObj, newObj)
			ex := newObj.(*v1.BGPAsNumber)
			fmt.Printf("AsNumber", ex.Spec.AsNumber, "Enable", ex.Spec.Enable)
			newex := ex.DeepCopy()
			newex.Status.Message = "Processed in handler"
			newex.Status.State = "Created"
			newex.Spec.AsNumber = "555"
			fmt.Printf("newex %+v\n", newex)
			if _, e := clientset.BgpV1().BGPAsNumbers("default").Update(newex); e != nil {
				fmt.Println("update satus error", e)
			}
		},
	},
	)
	stop := make(chan struct{})
	go bgpInformerFactory.Start(stop)

	// Wait forever
	select {}
}
