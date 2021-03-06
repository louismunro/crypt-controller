/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/bluehoodie/crypt-controller/pkg/apis/crypt/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CryptLister helps list Crypts.
type CryptLister interface {
	// List lists all Crypts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Crypt, err error)
	// Crypts returns an object that can list and get Crypts.
	Crypts(namespace string) CryptNamespaceLister
	CryptListerExpansion
}

// cryptLister implements the CryptLister interface.
type cryptLister struct {
	indexer cache.Indexer
}

// NewCryptLister returns a new CryptLister.
func NewCryptLister(indexer cache.Indexer) CryptLister {
	return &cryptLister{indexer: indexer}
}

// List lists all Crypts in the indexer.
func (s *cryptLister) List(selector labels.Selector) (ret []*v1alpha1.Crypt, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Crypt))
	})
	return ret, err
}

// Crypts returns an object that can list and get Crypts.
func (s *cryptLister) Crypts(namespace string) CryptNamespaceLister {
	return cryptNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CryptNamespaceLister helps list and get Crypts.
type CryptNamespaceLister interface {
	// List lists all Crypts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Crypt, err error)
	// Get retrieves the Crypt from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Crypt, error)
	CryptNamespaceListerExpansion
}

// cryptNamespaceLister implements the CryptNamespaceLister
// interface.
type cryptNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Crypts in the indexer for a given namespace.
func (s cryptNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Crypt, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Crypt))
	})
	return ret, err
}

// Get retrieves the Crypt from the indexer for a given namespace and name.
func (s cryptNamespaceLister) Get(name string) (*v1alpha1.Crypt, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("crypt"), name)
	}
	return obj.(*v1alpha1.Crypt), nil
}
