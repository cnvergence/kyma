// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	scheme "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UsageKindsGetter has a method to return a UsageKindInterface.
// A group's client should implement this interface.
type UsageKindsGetter interface {
	UsageKinds() UsageKindInterface
}

// UsageKindInterface has methods to work with UsageKind resources.
type UsageKindInterface interface {
	Create(ctx context.Context, usageKind *v1alpha1.UsageKind, opts v1.CreateOptions) (*v1alpha1.UsageKind, error)
	Update(ctx context.Context, usageKind *v1alpha1.UsageKind, opts v1.UpdateOptions) (*v1alpha1.UsageKind, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.UsageKind, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.UsageKindList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UsageKind, err error)
	UsageKindExpansion
}

// usageKinds implements UsageKindInterface
type usageKinds struct {
	client rest.Interface
}

// newUsageKinds returns a UsageKinds
func newUsageKinds(c *ServicecatalogV1alpha1Client) *usageKinds {
	return &usageKinds{
		client: c.RESTClient(),
	}
}

// Get takes name of the usageKind, and returns the corresponding usageKind object, and an error if there is any.
func (c *usageKinds) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UsageKind, err error) {
	result = &v1alpha1.UsageKind{}
	err = c.client.Get().
		Resource("usagekinds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UsageKinds that match those selectors.
func (c *usageKinds) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UsageKindList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UsageKindList{}
	err = c.client.Get().
		Resource("usagekinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested usageKinds.
func (c *usageKinds) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("usagekinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a usageKind and creates it.  Returns the server's representation of the usageKind, and an error, if there is any.
func (c *usageKinds) Create(ctx context.Context, usageKind *v1alpha1.UsageKind, opts v1.CreateOptions) (result *v1alpha1.UsageKind, err error) {
	result = &v1alpha1.UsageKind{}
	err = c.client.Post().
		Resource("usagekinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(usageKind).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a usageKind and updates it. Returns the server's representation of the usageKind, and an error, if there is any.
func (c *usageKinds) Update(ctx context.Context, usageKind *v1alpha1.UsageKind, opts v1.UpdateOptions) (result *v1alpha1.UsageKind, err error) {
	result = &v1alpha1.UsageKind{}
	err = c.client.Put().
		Resource("usagekinds").
		Name(usageKind.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(usageKind).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the usageKind and deletes it. Returns an error if one occurs.
func (c *usageKinds) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("usagekinds").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *usageKinds) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("usagekinds").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched usageKind.
func (c *usageKinds) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UsageKind, err error) {
	result = &v1alpha1.UsageKind{}
	err = c.client.Patch(pt).
		Resource("usagekinds").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}