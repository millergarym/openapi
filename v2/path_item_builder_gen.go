package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T20:40:48+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// PathItemBuilder is used to build an instance of PathItem. The user must
// call `Do()` after providing all the necessary information to
// build an instance of PathItem
type PathItemBuilder struct {
	target *pathItem
}

// Do finalizes the building process for PathItem and returns the result
func (b *PathItemBuilder) Do() (PathItem, error) {
	if err := b.target.Validate(); err != nil {
		return nil, errors.Wrap(err, `validation failed`)
	}
	return b.target, nil
}

// NewPathItem creates a new builder object for PathItem
func NewPathItem() *PathItemBuilder {
	return &PathItemBuilder{
		target: &pathItem{},
	}
}

// Parameters sets the Parameters field for object PathItem.
func (b *PathItemBuilder) Parameters(v ...Parameter) *PathItemBuilder {
	b.target.parameters = v
	return b
}

// Reference sets the $ref (reference) field for object PathItem.
func (b *PathItemBuilder) Reference(v string) *PathItemBuilder {
	b.target.reference = v
	return b
}