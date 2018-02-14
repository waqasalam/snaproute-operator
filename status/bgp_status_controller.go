package main

import (
	"flag"
	"fmt"
	apiextcs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"snaproute-operator/crd"
	"snaproute-operator/statusclient"
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
	fmt.Println("Start Status Server")
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
	err = crd.CreateCRD(clientset, "bgproute")
	if err != nil {
		panic(err)
	}

	// Create a new clientset which include our CRD schema
	crdcs, scheme, err := crd.NewClient(config)
	if err != nil {
		panic(err)
	}

	// Create a CRD client interface
	crdclient := client.CrdClient(crdcs, scheme, "default", crd.Crds["bgproute"].CrdPlural)

	// Create a new BGPAsNumber object and write to k8s you can use
	// bgp.yaml
	createfn := func() *crd.BGPRoute {
		return &crd.BGPRoute{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:   "bgproute-1.1.1.1-24",
				Labels: map[string]string{"mylabel": "test"},
			},
			Spec: crd.BGPRouteSpec{
				Prefix:  "1.1.1.1",
				Length:  24,
				Counter: 1,
			},
			Status: crd.BGPRouteStatus{
				State:   "created",
				Message: "Status crd",
			},
		}
	}

	bgproute := createfn()
	result, err := crdclient.Create(bgproute)
	if err == nil {
		fmt.Printf("CREATED: %#v\n", result)
	} else if apierrors.IsAlreadyExists(err) {
		fmt.Printf("ALREADY-- EXISTS: %#v\n", result)
	} else {
		panic(err)
	}
	state := "created"
	timer := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("timer fired", state)
			switch state {
			case "created":
				route, err := crdclient.Get(bgproute.Name, meta_v1.GetOptions{})
				if err != nil {
					fmt.Println("BGP route get failed", err)
				} else {
					route.Spec.Counter += 1
					if _, e := crdclient.Update(route); e != nil {
						fmt.Println("BGP route not updated")
					}
					state = "updated"
				}
			case "updated":

				err = crdclient.Delete(bgproute.Name, nil)
				if err != nil {
					fmt.Println("Delete error", err)
				}
				//delete

				state = "deleted"
			case "deleted":
				result, err := crdclient.Create(bgproute)
				if err == nil {
					fmt.Printf("CREATED: %#v\n", result)
				} else if apierrors.IsAlreadyExists(err) {
					fmt.Printf("ALREADY-- EXISTS: %#v\n", result)
				} else {
					panic(err)
				}

				if err != nil {
					fmt.Println("create failed", err)
				}

				state = "created"
			}
		}
	}
}
