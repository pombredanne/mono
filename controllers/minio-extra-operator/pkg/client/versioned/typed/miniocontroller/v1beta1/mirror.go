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

package v1beta1

import (
	"time"

	scheme "github.com/f110/wing/controllers/minio-extra-operator/pkg/client/versioned/scheme"
	v1beta1 "github.com/minio/minio-operator/pkg/apis/miniocontroller/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MirrorsGetter has a method to return a MirrorInterface.
// A group's client should implement this interface.
type MirrorsGetter interface {
	Mirrors(namespace string) MirrorInterface
}

// MirrorInterface has methods to work with Mirror resources.
type MirrorInterface interface {
	Create(*v1beta1.Mirror) (*v1beta1.Mirror, error)
	Update(*v1beta1.Mirror) (*v1beta1.Mirror, error)
	UpdateStatus(*v1beta1.Mirror) (*v1beta1.Mirror, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Mirror, error)
	List(opts v1.ListOptions) (*v1beta1.MirrorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Mirror, err error)
	MirrorExpansion
}

// mirrors implements MirrorInterface
type mirrors struct {
	client rest.Interface
	ns     string
}

// newMirrors returns a Mirrors
func newMirrors(c *MinV1beta1Client, namespace string) *mirrors {
	return &mirrors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mirror, and returns the corresponding mirror object, and an error if there is any.
func (c *mirrors) Get(name string, options v1.GetOptions) (result *v1beta1.Mirror, err error) {
	result = &v1beta1.Mirror{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mirrors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Mirrors that match those selectors.
func (c *mirrors) List(opts v1.ListOptions) (result *v1beta1.MirrorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.MirrorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mirrors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mirrors.
func (c *mirrors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mirrors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mirror and creates it.  Returns the server's representation of the mirror, and an error, if there is any.
func (c *mirrors) Create(mirror *v1beta1.Mirror) (result *v1beta1.Mirror, err error) {
	result = &v1beta1.Mirror{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mirrors").
		Body(mirror).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mirror and updates it. Returns the server's representation of the mirror, and an error, if there is any.
func (c *mirrors) Update(mirror *v1beta1.Mirror) (result *v1beta1.Mirror, err error) {
	result = &v1beta1.Mirror{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mirrors").
		Name(mirror.Name).
		Body(mirror).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mirrors) UpdateStatus(mirror *v1beta1.Mirror) (result *v1beta1.Mirror, err error) {
	result = &v1beta1.Mirror{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mirrors").
		Name(mirror.Name).
		SubResource("status").
		Body(mirror).
		Do().
		Into(result)
	return
}

// Delete takes name of the mirror and deletes it. Returns an error if one occurs.
func (c *mirrors) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mirrors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mirrors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mirrors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mirror.
func (c *mirrors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Mirror, err error) {
	result = &v1beta1.Mirror{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mirrors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
