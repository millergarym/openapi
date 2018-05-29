package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T20:40:48+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"log"
)

var _ = log.Printf

// PathsMutator is used to build an instance of Paths. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Paths with new values
type PathsMutator struct {
	proxy  *paths
	target *paths
}

// Do finalizes the matuation process for Paths and returns the result
func (b *PathsMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutatePaths creates a new mutator object for Paths
func MutatePaths(v Paths) *PathsMutator {
	return &PathsMutator{
		target: v.(*paths),
		proxy:  v.Clone().(*paths),
	}
}
func (b *PathsMutator) Extension(name string, value interface{}) *PathsMutator {
	if b.proxy.extensions == nil {
		b.proxy.extensions = Extensions{}
	}
	b.proxy.extensions[name] = value
	return b
}