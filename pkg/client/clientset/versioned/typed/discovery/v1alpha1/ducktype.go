/*
Copyright 2020 The Knative Authors

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
	"time"

	v1alpha1 "github.com/n3wscott/discovery/pkg/apis/discovery/v1alpha1"
	scheme "github.com/n3wscott/discovery/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DuckTypesGetter has a method to return a DuckTypeInterface.
// A group's client should implement this interface.
type DuckTypesGetter interface {
	DuckTypes() DuckTypeInterface
}

// DuckTypeInterface has methods to work with DuckType resources.
type DuckTypeInterface interface {
	Create(*v1alpha1.DuckType) (*v1alpha1.DuckType, error)
	Update(*v1alpha1.DuckType) (*v1alpha1.DuckType, error)
	UpdateStatus(*v1alpha1.DuckType) (*v1alpha1.DuckType, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DuckType, error)
	List(opts v1.ListOptions) (*v1alpha1.DuckTypeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DuckType, err error)
	DuckTypeExpansion
}

// duckTypes implements DuckTypeInterface
type duckTypes struct {
	client rest.Interface
}

// newDuckTypes returns a DuckTypes
func newDuckTypes(c *DiscoveryV1alpha1Client) *duckTypes {
	return &duckTypes{
		client: c.RESTClient(),
	}
}

// Get takes name of the duckType, and returns the corresponding duckType object, and an error if there is any.
func (c *duckTypes) Get(name string, options v1.GetOptions) (result *v1alpha1.DuckType, err error) {
	result = &v1alpha1.DuckType{}
	err = c.client.Get().
		Resource("ducktypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DuckTypes that match those selectors.
func (c *duckTypes) List(opts v1.ListOptions) (result *v1alpha1.DuckTypeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DuckTypeList{}
	err = c.client.Get().
		Resource("ducktypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested duckTypes.
func (c *duckTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ducktypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a duckType and creates it.  Returns the server's representation of the duckType, and an error, if there is any.
func (c *duckTypes) Create(duckType *v1alpha1.DuckType) (result *v1alpha1.DuckType, err error) {
	result = &v1alpha1.DuckType{}
	err = c.client.Post().
		Resource("ducktypes").
		Body(duckType).
		Do().
		Into(result)
	return
}

// Update takes the representation of a duckType and updates it. Returns the server's representation of the duckType, and an error, if there is any.
func (c *duckTypes) Update(duckType *v1alpha1.DuckType) (result *v1alpha1.DuckType, err error) {
	result = &v1alpha1.DuckType{}
	err = c.client.Put().
		Resource("ducktypes").
		Name(duckType.Name).
		Body(duckType).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *duckTypes) UpdateStatus(duckType *v1alpha1.DuckType) (result *v1alpha1.DuckType, err error) {
	result = &v1alpha1.DuckType{}
	err = c.client.Put().
		Resource("ducktypes").
		Name(duckType.Name).
		SubResource("status").
		Body(duckType).
		Do().
		Into(result)
	return
}

// Delete takes name of the duckType and deletes it. Returns an error if one occurs.
func (c *duckTypes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ducktypes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *duckTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ducktypes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched duckType.
func (c *duckTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DuckType, err error) {
	result = &v1alpha1.DuckType{}
	err = c.client.Patch(pt).
		Resource("ducktypes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
