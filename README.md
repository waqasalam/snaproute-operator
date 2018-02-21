# SnapRoute Custom Resources (CRD) Tutorial

## Build/Installtion
```
  make container 
  kubectl run bgp --image=bgp-crd:latest --image-pull-policy=Never
```
```
 kubectl create -f bgp.yaml
```



```go
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
```


