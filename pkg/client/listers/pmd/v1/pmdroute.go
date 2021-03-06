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

// This file was automatically generated by lister-gen

package v1

import (
	v1 "snaproute-operator/pkg/apis/pmd/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PMDRouteLister helps list PMDRoutes.
type PMDRouteLister interface {
	// List lists all PMDRoutes in the indexer.
	List(selector labels.Selector) (ret []*v1.PMDRoute, err error)
	// PMDRoutes returns an object that can list and get PMDRoutes.
	PMDRoutes(namespace string) PMDRouteNamespaceLister
	PMDRouteListerExpansion
}

// pMDRouteLister implements the PMDRouteLister interface.
type pMDRouteLister struct {
	indexer cache.Indexer
}

// NewPMDRouteLister returns a new PMDRouteLister.
func NewPMDRouteLister(indexer cache.Indexer) PMDRouteLister {
	return &pMDRouteLister{indexer: indexer}
}

// List lists all PMDRoutes in the indexer.
func (s *pMDRouteLister) List(selector labels.Selector) (ret []*v1.PMDRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PMDRoute))
	})
	return ret, err
}

// PMDRoutes returns an object that can list and get PMDRoutes.
func (s *pMDRouteLister) PMDRoutes(namespace string) PMDRouteNamespaceLister {
	return pMDRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PMDRouteNamespaceLister helps list and get PMDRoutes.
type PMDRouteNamespaceLister interface {
	// List lists all PMDRoutes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.PMDRoute, err error)
	// Get retrieves the PMDRoute from the indexer for a given namespace and name.
	Get(name string) (*v1.PMDRoute, error)
	PMDRouteNamespaceListerExpansion
}

// pMDRouteNamespaceLister implements the PMDRouteNamespaceLister
// interface.
type pMDRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PMDRoutes in the indexer for a given namespace.
func (s pMDRouteNamespaceLister) List(selector labels.Selector) (ret []*v1.PMDRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PMDRoute))
	})
	return ret, err
}

// Get retrieves the PMDRoute from the indexer for a given namespace and name.
func (s pMDRouteNamespaceLister) Get(name string) (*v1.PMDRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("pmdroute"), name)
	}
	return obj.(*v1.PMDRoute), nil
}
