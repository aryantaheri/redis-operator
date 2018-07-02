/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/spotahome/kooper/test/integration/operator/apis/superhero/v1alpha1"
	scheme "github.com/spotahome/kooper/test/integration/operator/client/k8s/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SpidermansGetter has a method to return a SpidermanInterface.
// A group's client should implement this interface.
type SpidermansGetter interface {
	Spidermans(namespace string) SpidermanInterface
}

// SpidermanInterface has methods to work with Spiderman resources.
type SpidermanInterface interface {
	Create(*v1alpha1.Spiderman) (*v1alpha1.Spiderman, error)
	Update(*v1alpha1.Spiderman) (*v1alpha1.Spiderman, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Spiderman, error)
	List(opts v1.ListOptions) (*v1alpha1.SpidermanList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Spiderman, err error)
	SpidermanExpansion
}

// spidermans implements SpidermanInterface
type spidermans struct {
	client rest.Interface
	ns     string
}

// newSpidermans returns a Spidermans
func newSpidermans(c *SuperheroV1alpha1Client, namespace string) *spidermans {
	return &spidermans{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the spiderman, and returns the corresponding spiderman object, and an error if there is any.
func (c *spidermans) Get(name string, options v1.GetOptions) (result *v1alpha1.Spiderman, err error) {
	result = &v1alpha1.Spiderman{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spidermans").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Spidermans that match those selectors.
func (c *spidermans) List(opts v1.ListOptions) (result *v1alpha1.SpidermanList, err error) {
	result = &v1alpha1.SpidermanList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spidermans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spidermans.
func (c *spidermans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("spidermans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a spiderman and creates it.  Returns the server's representation of the spiderman, and an error, if there is any.
func (c *spidermans) Create(spiderman *v1alpha1.Spiderman) (result *v1alpha1.Spiderman, err error) {
	result = &v1alpha1.Spiderman{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("spidermans").
		Body(spiderman).
		Do().
		Into(result)
	return
}

// Update takes the representation of a spiderman and updates it. Returns the server's representation of the spiderman, and an error, if there is any.
func (c *spidermans) Update(spiderman *v1alpha1.Spiderman) (result *v1alpha1.Spiderman, err error) {
	result = &v1alpha1.Spiderman{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("spidermans").
		Name(spiderman.Name).
		Body(spiderman).
		Do().
		Into(result)
	return
}

// Delete takes name of the spiderman and deletes it. Returns an error if one occurs.
func (c *spidermans) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spidermans").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spidermans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spidermans").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched spiderman.
func (c *spidermans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Spiderman, err error) {
	result = &v1alpha1.Spiderman{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("spidermans").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
