!#/bin/bash

~/kube/bin/deepcopy-gen --logtostderr --v=4 --input-dirs snaproute-operator/crd --go-header-file boilerplate.go.txt -O zz_generated
