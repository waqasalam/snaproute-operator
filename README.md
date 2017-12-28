# SnapRoute Custom Resources (CRD) Tutorial



## kube-crd

kube-crd demonstrates the CRD usage, it shoes how to:

1. Connect to the Kubernetes cluster 
2. Create the new CRD if it doesn't exist  
3. Create a new custom client 
4. Create a new Example object using the client library we created 
5. Create a controller that listens to events associated with new resources

The example CRD is in the following structure:


```go
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
```


