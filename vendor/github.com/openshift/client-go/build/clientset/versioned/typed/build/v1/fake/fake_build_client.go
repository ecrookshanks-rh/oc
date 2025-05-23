// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBuildV1 struct {
	*testing.Fake
}

func (c *FakeBuildV1) Builds(namespace string) v1.BuildInterface {
	return newFakeBuilds(c, namespace)
}

func (c *FakeBuildV1) BuildConfigs(namespace string) v1.BuildConfigInterface {
	return newFakeBuildConfigs(c, namespace)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBuildV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
