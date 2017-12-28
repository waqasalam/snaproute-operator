# SnapRoute Custom Resources (CRD) Tutorial

## Build/Installtion
```
  make container 
  kubectl run bgp --image=bgp-crd:latest --image-pull-policy=Never
```
###  Controller will startup and do following
```
  Connect to the Kubernetes cluster 
  Create the new CRD if it doesn't exist  
  Create a new custom client 
  Create a new BGPAsNumber object using the client library we created 
  Create a controller that listens to events associated with new CRD
```
###  Create BGPAsNumber object using yaml.
```
 kubectl create -f bgp.yaml
```

The BGPAsNumber CRD is in the following structure:


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


