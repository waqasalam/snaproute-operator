package v1

import (
	"fmt"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextcs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/wait"
	_ "k8s.io/client-go/rest"
	"reflect"
	"time"
)

const (
	CRDPlural   string = "pmdasnumbers"
	CRDGroup    string = "snaproute.com"
	CRDVersion  string = "v1"
	FullCRDName string = CRDPlural + "." + CRDGroup
)

// Create the CRD resource, ignore error if it already exists
func CreateCRD(clientset apiextcs.Interface) error {
	fmt.Println("I'm creating the PMD")
	crd := &apiextv1beta1.CustomResourceDefinition{
		ObjectMeta: meta_v1.ObjectMeta{Name: FullCRDName},
		Spec: apiextv1beta1.CustomResourceDefinitionSpec{
			Group:   CRDGroup,
			Version: CRDVersion,
			Scope:   apiextv1beta1.NamespaceScoped,
			Names: apiextv1beta1.CustomResourceDefinitionNames{
				Plural: CRDPlural,
				Kind:   reflect.TypeOf(PMDAsNumber{}).Name(),
			},
		},
	}

	_, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
	if err != nil && apierrors.IsAlreadyExists(err) {
		fmt.Println("CRD already exists")
		return nil
	}

	// Wait for the CRD to be created before we use it
	err = wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
		crd, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(FullCRDName, meta_v1.GetOptions{})
		if err != nil {
			fmt.Println("panic in wait")
			panic(err.Error())
		}

		fmt.Println("crd in wait", crd)
		for _, cond := range crd.Status.Conditions {
			switch cond.Type {
			case apiextv1beta1.Established:
				if cond.Status == apiextv1beta1.ConditionTrue {
					fmt.Printf("success already created no wait: %v\n", cond.Status)

					return true, err
				}
			case apiextv1beta1.NamesAccepted:
				if cond.Status == apiextv1beta1.ConditionFalse {
					fmt.Printf("Name conflict: %v\n", cond.Reason)
					fmt.Printf("error", err)
				}
			}
		}
		panic(err.Error())
	})
	return err
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient
// +genclient:noStatus

// Definition of our CRD BGPAsNumber class
type PMDAsNumber struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               PMDAsNumberSpec   `json:"spec"`
	Status             PMDAsNumberStatus `json:"status,omitempty"`
}
type PMDAsNumberSpec struct {
	AsNumber string `json:"asnumber"`
	Enable   bool   `json:"enable"`
}

type PMDAsNumberStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// k8s List Type
type PMDAsNumberList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []PMDAsNumber `json:"items"`
}
