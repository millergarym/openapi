package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T20:40:48+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"log"
)

var _ = log.Printf

// PathItemMutator is used to build an instance of PathItem. The user must
// call `Do()` after providing all the necessary information to
// the new instance of PathItem with new values
type PathItemMutator struct {
	proxy  *pathItem
	target *pathItem
}

// Do finalizes the matuation process for PathItem and returns the result
func (b *PathItemMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutatePathItem creates a new mutator object for PathItem
func MutatePathItem(v PathItem) *PathItemMutator {
	return &PathItemMutator{
		target: v.(*pathItem),
		proxy:  v.Clone().(*pathItem),
	}
}

func (b *PathItemMutator) ClearParameters() *PathItemMutator {
	b.proxy.parameters.Clear()
	return b
}

func (b *PathItemMutator) Parameter(value Parameter) *PathItemMutator {
	b.proxy.parameters = append(b.proxy.parameters, value)
	return b
}
func (b *PathItemMutator) Extension(name string, value interface{}) *PathItemMutator {
	if b.proxy.extensions == nil {
		b.proxy.extensions = Extensions{}
	}
	b.proxy.extensions[name] = value
	return b
}