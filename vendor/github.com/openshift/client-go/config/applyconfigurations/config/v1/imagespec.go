// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
)

// ImageSpecApplyConfiguration represents a declarative configuration of the ImageSpec type for use
// with apply.
type ImageSpecApplyConfiguration struct {
	AllowedRegistriesForImport []RegistryLocationApplyConfiguration      `json:"allowedRegistriesForImport,omitempty"`
	ExternalRegistryHostnames  []string                                  `json:"externalRegistryHostnames,omitempty"`
	AdditionalTrustedCA        *ConfigMapNameReferenceApplyConfiguration `json:"additionalTrustedCA,omitempty"`
	RegistrySources            *RegistrySourcesApplyConfiguration        `json:"registrySources,omitempty"`
	ImageStreamImportMode      *configv1.ImportModeType                  `json:"imageStreamImportMode,omitempty"`
}

// ImageSpecApplyConfiguration constructs a declarative configuration of the ImageSpec type for use with
// apply.
func ImageSpec() *ImageSpecApplyConfiguration {
	return &ImageSpecApplyConfiguration{}
}

// WithAllowedRegistriesForImport adds the given value to the AllowedRegistriesForImport field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllowedRegistriesForImport field.
func (b *ImageSpecApplyConfiguration) WithAllowedRegistriesForImport(values ...*RegistryLocationApplyConfiguration) *ImageSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAllowedRegistriesForImport")
		}
		b.AllowedRegistriesForImport = append(b.AllowedRegistriesForImport, *values[i])
	}
	return b
}

// WithExternalRegistryHostnames adds the given value to the ExternalRegistryHostnames field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExternalRegistryHostnames field.
func (b *ImageSpecApplyConfiguration) WithExternalRegistryHostnames(values ...string) *ImageSpecApplyConfiguration {
	for i := range values {
		b.ExternalRegistryHostnames = append(b.ExternalRegistryHostnames, values[i])
	}
	return b
}

// WithAdditionalTrustedCA sets the AdditionalTrustedCA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdditionalTrustedCA field is set to the value of the last call.
func (b *ImageSpecApplyConfiguration) WithAdditionalTrustedCA(value *ConfigMapNameReferenceApplyConfiguration) *ImageSpecApplyConfiguration {
	b.AdditionalTrustedCA = value
	return b
}

// WithRegistrySources sets the RegistrySources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RegistrySources field is set to the value of the last call.
func (b *ImageSpecApplyConfiguration) WithRegistrySources(value *RegistrySourcesApplyConfiguration) *ImageSpecApplyConfiguration {
	b.RegistrySources = value
	return b
}

// WithImageStreamImportMode sets the ImageStreamImportMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageStreamImportMode field is set to the value of the last call.
func (b *ImageSpecApplyConfiguration) WithImageStreamImportMode(value configv1.ImportModeType) *ImageSpecApplyConfiguration {
	b.ImageStreamImportMode = &value
	return b
}
