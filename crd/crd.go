package crd

import (
	"fmt"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextcs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"reflect"
	"time"
)

var Crds = map[string]struct {
	CrdFullName string
	CrdPlural   string
	CrdKind     string
}{
	"bgpasnumber": {"bgpasnumbers" + "." + CRDGroup, "bgapasnumbers", reflect.TypeOf(BGPAsNumber{}).Name()},
	"bgproute":    {"bgproutes" + "." + CRDGroup, "bgproutes", reflect.TypeOf(BGPRoute{}).Name()},
}

const (
	CRDPlural   string = "bgpasnumbers"
	CRDGroup    string = "snaproute.com"
	CRDVersion  string = "v1"
	FullCRDName string = CRDPlural + "." + CRDGroup
)

// Create the CRD resource, ignore error if it already exists
func CreateCRD(clientset apiextcs.Interface, crdname string) error {
	fmt.Println("I'm creating the CRD", Crds[crdname])
	crd := &apiextv1beta1.CustomResourceDefinition{
		ObjectMeta: meta_v1.ObjectMeta{Name: Crds[crdname].CrdFullName},
		Spec: apiextv1beta1.CustomResourceDefinitionSpec{
			Group:   CRDGroup,
			Version: CRDVersion,
			Scope:   apiextv1beta1.NamespaceScoped,
			Names: apiextv1beta1.CustomResourceDefinitionNames{
				Plural: Crds[crdname].CrdPlural,
				Kind:   Crds[crdname].CrdKind,
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

// Definition of our CRD BGPAsNumber class
type BGPAsNumber struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               BGPAsNumberSpec   `json:"spec"`
	Status             BGPAsNumberStatus `json:"status,omitempty"`
}
type BGPAsNumberSpec struct {
	AsNumber string `json:"asnumber"`
	Enable   bool   `json:"enable"`
}

type BGPAsNumberStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// k8s List Type
type BGPAsNumberList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []BGPAsNumber `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Definition of our CRD BGPRoute class
type BGPRoute struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               BGPRouteSpec   `json:"spec"`
	Status             BGPRouteStatus `json:"status,omitempty"`
}
type BGPRouteSpec struct {
	Prefix  string `json:"prefix"`
	Length  uint32 `json:"length"`
	Counter uint32 `json:"counter"`
}

type BGPRouteStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// k8s List Type
type BGPRouteList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []BGPRoute `json:"items"`
}

// Create a  Rest client with the new CRD Schema
var SchemeGroupVersion = schema.GroupVersion{Group: CRDGroup, Version: CRDVersion}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&BGPAsNumber{},
		&BGPAsNumberList{},
		&BGPRoute{},
		&BGPRouteList{},
	)
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

func NewClient(cfg *rest.Config) (*rest.RESTClient, *runtime.Scheme, error) {
	scheme := runtime.NewScheme()
	SchemeBuilder := runtime.NewSchemeBuilder(addKnownTypes)
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		return nil, nil, err
	}
	config := *cfg
	config.GroupVersion = &SchemeGroupVersion
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.DirectCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme)}

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, nil, err
	}
	return client, scheme, nil
}
