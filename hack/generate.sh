!#/bin/bash

~/kubernetes-1.9.2/staging/bin/deepcopy-gen --logtostderr --v=4 --input-dirs snaproute-operator/crd --go-header-file boilerplate.go.txt -O zz_generated
