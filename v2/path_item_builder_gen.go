package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
	"sync"
)

var _ = errors.Cause

// PathItemBuilder is used to build an instance of PathItem. The user must
// call `Build()` after providing all the necessary information to
// build an instance of PathItem.
// Builders may NOT be reused. It must be created for every instance
// of PathItem that you want to create
type PathItemBuilder struct {
	lock   sync.Locker
	target *pathItem
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *PathItemBuilder) MustBuild(options ...Option) PathItem {
	v, err := b.Build(options...)
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for PathItem and returns the result
// By default, Build() will validate if the given structure is valid
func (b *PathItemBuilder) Build(options ...Option) (PathItem, error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return nil, errors.New(`builder has already been used`)
	}
	validate := true
	for _, option := range options {
		switch option.Name() {
		case optkeyValidate:
			validate = option.Value().(bool)
		}
	}
	if validate {
		if err := b.target.Validate(false); err != nil {
			return nil, errors.Wrap(err, `validation failed`)
		}
	}
	defer func() { b.target = nil }()
	return b.target, nil
}

// NewPathItem creates a new builder object for PathItem
func NewPathItem(options ...Option) *PathItemBuilder {
	var lock sync.Locker = &sync.Mutex{}
	for _, option := range options {
		switch option.Name() {
		case optkeyLocker:
			lock = option.Value().(sync.Locker)
		}
	}
	var b PathItemBuilder
	if lock == nil {
		lock = nilLock{}
	}
	b.lock = lock
	b.target = &pathItem{}
	return &b
}

// Parameters sets the parameters field for object PathItem.
func (b *PathItemBuilder) Parameters(v ...Parameter) *PathItemBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.parameters = v
	return b
}

// Reference sets the $ref (reference) field for object PathItem.
func (b *PathItemBuilder) Reference(v string) *PathItemBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.reference = v
	return b
}

// Extension sets an arbitrary element (an extension) to the
// object PathItem. The extension name should start with a "x-"
func (b *PathItemBuilder) Extension(name string, value interface{}) *PathItemBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.extensions[name] = value
	return b
}
