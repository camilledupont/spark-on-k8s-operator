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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "k8s.io/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1alpha1"
)

// FakeSparkApplications implements SparkApplicationInterface
type FakeSparkApplications struct {
	Fake *FakeSparkoperatorV1alpha1
	ns   string
}

var sparkapplicationsResource = schema.GroupVersionResource{Group: "sparkoperator", Version: "v1alpha1", Resource: "sparkapplications"}

var sparkapplicationsKind = schema.GroupVersionKind{Group: "sparkoperator", Version: "v1alpha1", Kind: "SparkApplication"}

// Get takes name of the sparkApplication, and returns the corresponding sparkApplication object, and an error if there is any.
func (c *FakeSparkApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.SparkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sparkapplicationsResource, c.ns, name), &v1alpha1.SparkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SparkApplication), err
}

// List takes label and field selectors, and returns the list of SparkApplications that match those selectors.
func (c *FakeSparkApplications) List(opts v1.ListOptions) (result *v1alpha1.SparkApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sparkapplicationsResource, sparkapplicationsKind, c.ns, opts), &v1alpha1.SparkApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SparkApplicationList{}
	for _, item := range obj.(*v1alpha1.SparkApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sparkApplications.
func (c *FakeSparkApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sparkapplicationsResource, c.ns, opts))

}

// Create takes the representation of a sparkApplication and creates it.  Returns the server's representation of the sparkApplication, and an error, if there is any.
func (c *FakeSparkApplications) Create(sparkApplication *v1alpha1.SparkApplication) (result *v1alpha1.SparkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sparkapplicationsResource, c.ns, sparkApplication), &v1alpha1.SparkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SparkApplication), err
}

// Update takes the representation of a sparkApplication and updates it. Returns the server's representation of the sparkApplication, and an error, if there is any.
func (c *FakeSparkApplications) Update(sparkApplication *v1alpha1.SparkApplication) (result *v1alpha1.SparkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sparkapplicationsResource, c.ns, sparkApplication), &v1alpha1.SparkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SparkApplication), err
}

// Delete takes name of the sparkApplication and deletes it. Returns an error if one occurs.
func (c *FakeSparkApplications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sparkapplicationsResource, c.ns, name), &v1alpha1.SparkApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSparkApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sparkapplicationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SparkApplicationList{})
	return err
}

// Patch applies the patch and returns the patched sparkApplication.
func (c *FakeSparkApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SparkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sparkapplicationsResource, c.ns, name, data, subresources...), &v1alpha1.SparkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SparkApplication), err
}
