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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.io/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1alpha1"
	scheme "k8s.io/spark-on-k8s-operator/pkg/client/clientset/versioned/scheme"
)

// SparkApplicationsGetter has a method to return a SparkApplicationInterface.
// A group's client should implement this interface.
type SparkApplicationsGetter interface {
	SparkApplications(namespace string) SparkApplicationInterface
}

// SparkApplicationInterface has methods to work with SparkApplication resources.
type SparkApplicationInterface interface {
	Create(*v1alpha1.SparkApplication) (*v1alpha1.SparkApplication, error)
	Update(*v1alpha1.SparkApplication) (*v1alpha1.SparkApplication, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SparkApplication, error)
	List(opts v1.ListOptions) (*v1alpha1.SparkApplicationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SparkApplication, err error)
	SparkApplicationExpansion
}

// sparkApplications implements SparkApplicationInterface
type sparkApplications struct {
	client rest.Interface
	ns     string
}

// newSparkApplications returns a SparkApplications
func newSparkApplications(c *SparkoperatorV1alpha1Client, namespace string) *sparkApplications {
	return &sparkApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sparkApplication, and returns the corresponding sparkApplication object, and an error if there is any.
func (c *sparkApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.SparkApplication, err error) {
	result = &v1alpha1.SparkApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sparkapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SparkApplications that match those selectors.
func (c *sparkApplications) List(opts v1.ListOptions) (result *v1alpha1.SparkApplicationList, err error) {
	result = &v1alpha1.SparkApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sparkApplications.
func (c *sparkApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a sparkApplication and creates it.  Returns the server's representation of the sparkApplication, and an error, if there is any.
func (c *sparkApplications) Create(sparkApplication *v1alpha1.SparkApplication) (result *v1alpha1.SparkApplication, err error) {
	result = &v1alpha1.SparkApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sparkapplications").
		Body(sparkApplication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sparkApplication and updates it. Returns the server's representation of the sparkApplication, and an error, if there is any.
func (c *sparkApplications) Update(sparkApplication *v1alpha1.SparkApplication) (result *v1alpha1.SparkApplication, err error) {
	result = &v1alpha1.SparkApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sparkapplications").
		Name(sparkApplication.Name).
		Body(sparkApplication).
		Do().
		Into(result)
	return
}

// Delete takes name of the sparkApplication and deletes it. Returns an error if one occurs.
func (c *sparkApplications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sparkapplications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sparkApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sparkapplications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sparkApplication.
func (c *sparkApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SparkApplication, err error) {
	result = &v1alpha1.SparkApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sparkapplications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
