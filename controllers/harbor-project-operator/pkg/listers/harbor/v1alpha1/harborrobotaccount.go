/*
MIT License

Copyright (c) 2020 Fumihiro Ito

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/f110/wing/controllers/harbor-project-operator/pkg/api/harbor/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HarborRobotAccountLister helps list HarborRobotAccounts.
type HarborRobotAccountLister interface {
	// List lists all HarborRobotAccounts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.HarborRobotAccount, err error)
	// HarborRobotAccounts returns an object that can list and get HarborRobotAccounts.
	HarborRobotAccounts(namespace string) HarborRobotAccountNamespaceLister
	HarborRobotAccountListerExpansion
}

// harborRobotAccountLister implements the HarborRobotAccountLister interface.
type harborRobotAccountLister struct {
	indexer cache.Indexer
}

// NewHarborRobotAccountLister returns a new HarborRobotAccountLister.
func NewHarborRobotAccountLister(indexer cache.Indexer) HarborRobotAccountLister {
	return &harborRobotAccountLister{indexer: indexer}
}

// List lists all HarborRobotAccounts in the indexer.
func (s *harborRobotAccountLister) List(selector labels.Selector) (ret []*v1alpha1.HarborRobotAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HarborRobotAccount))
	})
	return ret, err
}

// HarborRobotAccounts returns an object that can list and get HarborRobotAccounts.
func (s *harborRobotAccountLister) HarborRobotAccounts(namespace string) HarborRobotAccountNamespaceLister {
	return harborRobotAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HarborRobotAccountNamespaceLister helps list and get HarborRobotAccounts.
type HarborRobotAccountNamespaceLister interface {
	// List lists all HarborRobotAccounts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.HarborRobotAccount, err error)
	// Get retrieves the HarborRobotAccount from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.HarborRobotAccount, error)
	HarborRobotAccountNamespaceListerExpansion
}

// harborRobotAccountNamespaceLister implements the HarborRobotAccountNamespaceLister
// interface.
type harborRobotAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HarborRobotAccounts in the indexer for a given namespace.
func (s harborRobotAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HarborRobotAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HarborRobotAccount))
	})
	return ret, err
}

// Get retrieves the HarborRobotAccount from the indexer for a given namespace and name.
func (s harborRobotAccountNamespaceLister) Get(name string) (*v1alpha1.HarborRobotAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("harborrobotaccount"), name)
	}
	return obj.(*v1alpha1.HarborRobotAccount), nil
}
