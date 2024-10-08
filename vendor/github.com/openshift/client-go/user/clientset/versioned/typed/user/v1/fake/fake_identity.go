// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/user/v1"
	userv1 "github.com/openshift/client-go/user/applyconfigurations/user/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIdentities implements IdentityInterface
type FakeIdentities struct {
	Fake *FakeUserV1
}

var identitiesResource = v1.SchemeGroupVersion.WithResource("identities")

var identitiesKind = v1.SchemeGroupVersion.WithKind("Identity")

// Get takes name of the identity, and returns the corresponding identity object, and an error if there is any.
func (c *FakeIdentities) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Identity, err error) {
	emptyResult := &v1.Identity{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(identitiesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Identity), err
}

// List takes label and field selectors, and returns the list of Identities that match those selectors.
func (c *FakeIdentities) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IdentityList, err error) {
	emptyResult := &v1.IdentityList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(identitiesResource, identitiesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.IdentityList{ListMeta: obj.(*v1.IdentityList).ListMeta}
	for _, item := range obj.(*v1.IdentityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identities.
func (c *FakeIdentities) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(identitiesResource, opts))
}

// Create takes the representation of a identity and creates it.  Returns the server's representation of the identity, and an error, if there is any.
func (c *FakeIdentities) Create(ctx context.Context, identity *v1.Identity, opts metav1.CreateOptions) (result *v1.Identity, err error) {
	emptyResult := &v1.Identity{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(identitiesResource, identity, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Identity), err
}

// Update takes the representation of a identity and updates it. Returns the server's representation of the identity, and an error, if there is any.
func (c *FakeIdentities) Update(ctx context.Context, identity *v1.Identity, opts metav1.UpdateOptions) (result *v1.Identity, err error) {
	emptyResult := &v1.Identity{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(identitiesResource, identity, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Identity), err
}

// Delete takes name of the identity and deletes it. Returns an error if one occurs.
func (c *FakeIdentities) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(identitiesResource, name, opts), &v1.Identity{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentities) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(identitiesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.IdentityList{})
	return err
}

// Patch applies the patch and returns the patched identity.
func (c *FakeIdentities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Identity, err error) {
	emptyResult := &v1.Identity{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(identitiesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Identity), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied identity.
func (c *FakeIdentities) Apply(ctx context.Context, identity *userv1.IdentityApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Identity, err error) {
	if identity == nil {
		return nil, fmt.Errorf("identity provided to Apply must not be nil")
	}
	data, err := json.Marshal(identity)
	if err != nil {
		return nil, err
	}
	name := identity.Name
	if name == nil {
		return nil, fmt.Errorf("identity.Name must be provided to Apply")
	}
	emptyResult := &v1.Identity{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(identitiesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Identity), err
}
