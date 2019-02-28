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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/yard-turkey/lib-bucket-provisioner/pkg/apis/objectbucket.io/v1alpha1"
	scheme "github.com/yard-turkey/lib-bucket-provisioner/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ObjectBucketClaimsGetter has a method to return a ObjectBucketClaimInterface.
// A group's client should implement this interface.
type ObjectBucketClaimsGetter interface {
	ObjectBucketClaims(namespace string) ObjectBucketClaimInterface
}

// ObjectBucketClaimInterface has methods to work with ObjectBucketClaim resources.
type ObjectBucketClaimInterface interface {
	Create(*v1alpha1.ObjectBucketClaim) (*v1alpha1.ObjectBucketClaim, error)
	Update(*v1alpha1.ObjectBucketClaim) (*v1alpha1.ObjectBucketClaim, error)
	UpdateStatus(*v1alpha1.ObjectBucketClaim) (*v1alpha1.ObjectBucketClaim, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ObjectBucketClaim, error)
	List(opts v1.ListOptions) (*v1alpha1.ObjectBucketClaimList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ObjectBucketClaim, err error)
	ObjectBucketClaimExpansion
}

// objectBucketClaims implements ObjectBucketClaimInterface
type objectBucketClaims struct {
	client rest.Interface
	ns     string
}

// newObjectBucketClaims returns a ObjectBucketClaims
func newObjectBucketClaims(c *ObjectbucketV1alpha1Client, namespace string) *objectBucketClaims {
	return &objectBucketClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the objectBucketClaim, and returns the corresponding objectBucketClaim object, and an error if there is any.
func (c *objectBucketClaims) Get(name string, options v1.GetOptions) (result *v1alpha1.ObjectBucketClaim, err error) {
	result = &v1alpha1.ObjectBucketClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("objectbucketclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ObjectBucketClaims that match those selectors.
func (c *objectBucketClaims) List(opts v1.ListOptions) (result *v1alpha1.ObjectBucketClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ObjectBucketClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("objectbucketclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested objectBucketClaims.
func (c *objectBucketClaims) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("objectbucketclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a objectBucketClaim and creates it.  Returns the server's representation of the objectBucketClaim, and an error, if there is any.
func (c *objectBucketClaims) Create(objectBucketClaim *v1alpha1.ObjectBucketClaim) (result *v1alpha1.ObjectBucketClaim, err error) {
	result = &v1alpha1.ObjectBucketClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("objectbucketclaims").
		Body(objectBucketClaim).
		Do().
		Into(result)
	return
}

// Update takes the representation of a objectBucketClaim and updates it. Returns the server's representation of the objectBucketClaim, and an error, if there is any.
func (c *objectBucketClaims) Update(objectBucketClaim *v1alpha1.ObjectBucketClaim) (result *v1alpha1.ObjectBucketClaim, err error) {
	result = &v1alpha1.ObjectBucketClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("objectbucketclaims").
		Name(objectBucketClaim.Name).
		Body(objectBucketClaim).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *objectBucketClaims) UpdateStatus(objectBucketClaim *v1alpha1.ObjectBucketClaim) (result *v1alpha1.ObjectBucketClaim, err error) {
	result = &v1alpha1.ObjectBucketClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("objectbucketclaims").
		Name(objectBucketClaim.Name).
		SubResource("status").
		Body(objectBucketClaim).
		Do().
		Into(result)
	return
}

// Delete takes name of the objectBucketClaim and deletes it. Returns an error if one occurs.
func (c *objectBucketClaims) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("objectbucketclaims").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *objectBucketClaims) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("objectbucketclaims").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched objectBucketClaim.
func (c *objectBucketClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ObjectBucketClaim, err error) {
	result = &v1alpha1.ObjectBucketClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("objectbucketclaims").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
