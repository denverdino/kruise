/*
Copyright 2020 The Kruise Authors.

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

	v1alpha1 "github.com/openkruise/kruise/apis/apps/v1alpha1"
	scheme "github.com/openkruise/kruise/pkg/client/clientset/versioned/scheme"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UnitedDeploymentsGetter has a method to return a UnitedDeploymentInterface.
// A group's client should implement this interface.
type UnitedDeploymentsGetter interface {
	UnitedDeployments(namespace string) UnitedDeploymentInterface
}

// UnitedDeploymentInterface has methods to work with UnitedDeployment resources.
type UnitedDeploymentInterface interface {
	Create(*v1alpha1.UnitedDeployment) (*v1alpha1.UnitedDeployment, error)
	Update(*v1alpha1.UnitedDeployment) (*v1alpha1.UnitedDeployment, error)
	UpdateStatus(*v1alpha1.UnitedDeployment) (*v1alpha1.UnitedDeployment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.UnitedDeployment, error)
	List(opts v1.ListOptions) (*v1alpha1.UnitedDeploymentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UnitedDeployment, err error)
	GetScale(unitedDeploymentName string, options v1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(unitedDeploymentName string, scale *autoscalingv1.Scale) (*autoscalingv1.Scale, error)

	UnitedDeploymentExpansion
}

// unitedDeployments implements UnitedDeploymentInterface
type unitedDeployments struct {
	client rest.Interface
	ns     string
}

// newUnitedDeployments returns a UnitedDeployments
func newUnitedDeployments(c *AppsV1alpha1Client, namespace string) *unitedDeployments {
	return &unitedDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the unitedDeployment, and returns the corresponding unitedDeployment object, and an error if there is any.
func (c *unitedDeployments) Get(name string, options v1.GetOptions) (result *v1alpha1.UnitedDeployment, err error) {
	result = &v1alpha1.UnitedDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UnitedDeployments that match those selectors.
func (c *unitedDeployments) List(opts v1.ListOptions) (result *v1alpha1.UnitedDeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UnitedDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("uniteddeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested unitedDeployments.
func (c *unitedDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("uniteddeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a unitedDeployment and creates it.  Returns the server's representation of the unitedDeployment, and an error, if there is any.
func (c *unitedDeployments) Create(unitedDeployment *v1alpha1.UnitedDeployment) (result *v1alpha1.UnitedDeployment, err error) {
	result = &v1alpha1.UnitedDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Body(unitedDeployment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a unitedDeployment and updates it. Returns the server's representation of the unitedDeployment, and an error, if there is any.
func (c *unitedDeployments) Update(unitedDeployment *v1alpha1.UnitedDeployment) (result *v1alpha1.UnitedDeployment, err error) {
	result = &v1alpha1.UnitedDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Name(unitedDeployment.Name).
		Body(unitedDeployment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *unitedDeployments) UpdateStatus(unitedDeployment *v1alpha1.UnitedDeployment) (result *v1alpha1.UnitedDeployment, err error) {
	result = &v1alpha1.UnitedDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Name(unitedDeployment.Name).
		SubResource("status").
		Body(unitedDeployment).
		Do().
		Into(result)
	return
}

// Delete takes name of the unitedDeployment and deletes it. Returns an error if one occurs.
func (c *unitedDeployments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *unitedDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("uniteddeployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched unitedDeployment.
func (c *unitedDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UnitedDeployment, err error) {
	result = &v1alpha1.UnitedDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("uniteddeployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

// GetScale takes name of the unitedDeployment, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *unitedDeployments) GetScale(unitedDeploymentName string, options v1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Name(unitedDeploymentName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *unitedDeployments) UpdateScale(unitedDeploymentName string, scale *autoscalingv1.Scale) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Name(unitedDeploymentName).
		SubResource("scale").
		Body(scale).
		Do().
		Into(result)
	return
}
