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
	"snaproute-operator/pkg/apis/pmd/v1"
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

	// create clientset
	clientset, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	clt, err := apiextcs.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//Create all the crds
	err = v1.CreateCRD(clt, "pmdasnumber")
	if err != nil {
		panic(err)
	}

	err = v1.CreateCRD(clt, "pmdroute")
	if err != nil {
		panic(err)
	}

	scheme.AddToScheme(scheme.Scheme)

	pmdInformerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	fmt.Println("create informers")
	asinformer := pmdInformerFactory.Pmd().V1().PMDAsNumbers()
	routeinformer := pmdInformerFactory.Pmd().V1().PMDRoutes()

	pmdasnumber := &v1.PMDAsNumber{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:   "pmdasnumber1",
			Labels: map[string]string{"mylabel": "test"},
		},
		Spec: v1.PMDAsNumberSpec{
			AsNumber: "one",
			Enable:   true,
		},
		Status: v1.PMDAsNumberStatus{
			State:   "created",
			Message: "Created, not -- processed yet",
		},
	}

	result, err := clientset.PmdV1().PMDAsNumbers("default").Create(pmdasnumber)
	if err == nil {
		fmt.Printf("created: %#v\n", result)
	} else if apierrors.IsAlreadyExists(err) {
		fmt.Printf("already created: %#v\n", result)
	} else {
		panic(err)
	}
	fmt.Printf("crd creation done")

	// List all BGP AsNumber objects
	items, err := clientset.PmdV1().PMDAsNumbers("default").List(meta_v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Do list operations :List:\n%s\n", items)
	fmt.Println("List operation done")
	// BGP Controller
	// Watch for changes in BGP objects and fire Add, Delete, Update callbacks

	asinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("add -- crd shared: %s \n", obj)
			ex := obj.(*v1.PMDAsNumber)
			newex := ex.DeepCopy()
			newex.Status.Message = "Processed in handler"
			newex.Status.State = "Created"
			newex.Spec.AsNumber = "555"
			if _, e := clientset.PmdV1().PMDAsNumbers("default").Update(newex); e != nil {
				fmt.Println("update satus error", e)
			}
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Printf("delete: %s shared \n", obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Printf("Update -- old: %s \n      New: %s\n", oldObj, newObj)
			ex := newObj.(*v1.PMDAsNumber)
			newex := ex.DeepCopy()
			newex.Status.Message = "Processed in handler"
			newex.Status.State = "Created"
			newex.Spec.AsNumber = "555"
			if _, e := clientset.PmdV1().PMDAsNumbers("default").Update(newex); e != nil {
				fmt.Println("update satus error", e)
			}
		},
	},
	)

	routeinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("add -- route: %s \n", obj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Printf("delete route: %s shared \n", obj)

		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Printf("Update route: %s \n      New: %s\n", oldObj, newObj)
		},
	},
	)
	stop := make(chan struct{})
	go pmdInformerFactory.Start(stop)

	// Wait forever
	select {}
}
