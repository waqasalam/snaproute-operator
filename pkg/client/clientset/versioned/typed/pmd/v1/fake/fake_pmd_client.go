/*
Copyright 2018 The Voyager Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package fake

import (
	v1 "snaproute-operator/pkg/client/clientset/versioned/typed/pmd/v1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePmdV1 struct {
	*testing.Fake
}

func (c *FakePmdV1) PMDAsNumbers(namespace string) v1.PMDAsNumberInterface {
	return &FakePMDAsNumbers{c, namespace}
}

func (c *FakePmdV1) PMDRoutes(namespace string) v1.PMDRouteInterface {
	return &FakePMDRoutes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePmdV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
