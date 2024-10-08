// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// IDRangeApplyConfiguration represents a declarative configuration of the IDRange type for use
// with apply.
type IDRangeApplyConfiguration struct {
	Min *int64 `json:"min,omitempty"`
	Max *int64 `json:"max,omitempty"`
}

// IDRangeApplyConfiguration constructs a declarative configuration of the IDRange type for use with
// apply.
func IDRange() *IDRangeApplyConfiguration {
	return &IDRangeApplyConfiguration{}
}

// WithMin sets the Min field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Min field is set to the value of the last call.
func (b *IDRangeApplyConfiguration) WithMin(value int64) *IDRangeApplyConfiguration {
	b.Min = &value
	return b
}

// WithMax sets the Max field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Max field is set to the value of the last call.
func (b *IDRangeApplyConfiguration) WithMax(value int64) *IDRangeApplyConfiguration {
	b.Max = &value
	return b
}
