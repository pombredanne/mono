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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "go.f110.dev/mono/go/pkg/k8s/client/versioned"
	grafanav1alpha1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/grafana/v1alpha1"
	fakegrafanav1alpha1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/grafana/v1alpha1/fake"
	harborv1alpha1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/harbor/v1alpha1"
	fakeharborv1alpha1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/harbor/v1alpha1/fake"
	miniov1alpha1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/minio/v1alpha1"
	fakeminiov1alpha1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/minio/v1alpha1/fake"
	miniocontrollerv1beta1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/miniocontroller/v1beta1"
	fakeminiocontrollerv1beta1 "go.f110.dev/mono/go/pkg/k8s/client/versioned/typed/miniocontroller/v1beta1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// GrafanaV1alpha1 retrieves the GrafanaV1alpha1Client
func (c *Clientset) GrafanaV1alpha1() grafanav1alpha1.GrafanaV1alpha1Interface {
	return &fakegrafanav1alpha1.FakeGrafanaV1alpha1{Fake: &c.Fake}
}

// HarborV1alpha1 retrieves the HarborV1alpha1Client
func (c *Clientset) HarborV1alpha1() harborv1alpha1.HarborV1alpha1Interface {
	return &fakeharborv1alpha1.FakeHarborV1alpha1{Fake: &c.Fake}
}

// MinioV1alpha1 retrieves the MinioV1alpha1Client
func (c *Clientset) MinioV1alpha1() miniov1alpha1.MinioV1alpha1Interface {
	return &fakeminiov1alpha1.FakeMinioV1alpha1{Fake: &c.Fake}
}

// MiniocontrollerV1beta1 retrieves the MiniocontrollerV1beta1Client
func (c *Clientset) MiniocontrollerV1beta1() miniocontrollerv1beta1.MiniocontrollerV1beta1Interface {
	return &fakeminiocontrollerv1beta1.FakeMiniocontrollerV1beta1{Fake: &c.Fake}
}
