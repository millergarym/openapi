package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"
	"sort"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = sort.Strings
var _ = errors.Cause

func (v *items) IsValid() bool {
	return v != nil
}

func (v *items) Type() PrimitiveType {
	return v.typ
}

func (v *items) Format() string {
	return v.format
}

func (v *items) Items() Items {
	return v.items
}

func (v *items) CollectionFormat() CollectionFormat {
	return v.collectionFormat
}

func (v *items) Default() interface{} {
	return v.defaultValue
}

// HasMaximum returns true if the value for maximum has been
// explicitly specified
func (v *items) HasMaximum() bool {
	return v.maximum != nil
}

// Maximum returns the value of maximum. If the value has not
// been explicitly, set, the zero value will be returned
func (v *items) Maximum() float64 {
	if !v.HasMaximum() {
		return 0
	}
	return *v.maximum
}

// HasExclusiveMaximum returns true if the value for exclusiveMaximum has been
// explicitly specified
func (v *items) HasExclusiveMaximum() bool {
	return v.exclusiveMaximum != nil
}

// ExclusiveMaximum returns the value of exclusiveMaximum. If the value has not
// been explicitly, set, the zero value will be returned
func (v *items) ExclusiveMaximum() float64 {
	if !v.HasExclusiveMaximum() {
		return 0
	}
	return *v.exclusiveMaximum
}

// HasMinimum returns true if the value for minimum has been
// explicitly specified
func (v *items) HasMinimum() bool {
	return v.minimum != nil
}

// Minimum returns the value of minimum. If the value has not
// been explicitly, set, the zero value will be returned
func (v *items) Minimum() float64 {
	if !v.HasMinimum() {
		return 0
	}
	return *v.minimum
}

// HasExclusiveMinimum returns true if the value for exclusiveMinimum has been
// explicitly specified
func (v *items) HasExclusiveMinimum() bool {
	return v.exclusiveMinimum != nil
}

// ExclusiveMinimum returns the value of exclusiveMinimum. If the value has not
// been explicitly, set, the zero value will be returned
func (v *items) ExclusiveMinimum() float64 {
	if !v.HasExclusiveMinimum() {
		return 0
	}
	return *v.exclusiveMinimum
}

// HasMaxLength returns true if the value for maxLength has been
// explicitly specified
func (v *items) HasMaxLength() bool {
	return v.maxLength != nil
}

// MaxLength returns the value of maxLength. If the value has not
// been explicitly, set, the zero value will be returned
func (v *items) MaxLength() int {
	if !v.HasMaxLength() {
		return 0
	}
	return *v.maxLength
}

// HasMinLength returns true if the value for minLength has been
// explicitly specified
func (v *items) HasMinLength() bool {
	return v.minLength != nil
}

// MinLength returns the value of minLength. If the value has not
// been explicitly, set, the zero value will be returned
func (v *items) MinLength() int {
	if !v.HasMinLength() {
		return 0
	}
	return *v.minLength
}

func (v *items) Pattern() string {
	return v.pattern
}

// HasMaxItems returns true if the value for maxItems has been
// explicitly specified
func (v *items) HasMaxItems() bool {
	return v.maxItems != nil
}

// MaxItems returns the value of maxItems. If the value has not
// been explicitly, set, the zero value will be returned
func (v *items) MaxItems() int {
	if !v.HasMaxItems() {
		return 0
	}
	return *v.maxItems
}

// HasMinItems returns true if the value for minItems has been
// explicitly specified
func (v *items) HasMinItems() bool {
	return v.minItems != nil
}

// MinItems returns the value of minItems. If the value has not
// been explicitly, set, the zero value will be returned
func (v *items) MinItems() int {
	if !v.HasMinItems() {
		return 0
	}
	return *v.minItems
}

// HasUniqueItems returns true if the value for uniqueItems has been
// explicitly specified
func (v *items) HasUniqueItems() bool {
	return v.uniqueItems != nil
}

// UniqueItems returns the value of uniqueItems. If the value has not
// been explicitly, set, the zero value will be returned
func (v *items) UniqueItems() bool {
	if !v.HasUniqueItems() {
		return false
	}
	return *v.uniqueItems
}

func (v *items) Enum() *InterfaceListIterator {
	var items []interface{}
	for _, item := range v.enum {
		items = append(items, item)
	}
	var iter InterfaceListIterator
	iter.size = len(items)
	iter.items = items
	return &iter
}

// HasMultipleOf returns true if the value for multipleOf has been
// explicitly specified
func (v *items) HasMultipleOf() bool {
	return v.multipleOf != nil
}

// MultipleOf returns the value of multipleOf. If the value has not
// been explicitly, set, the zero value will be returned
func (v *items) MultipleOf() float64 {
	if !v.HasMultipleOf() {
		return 0
	}
	return *v.multipleOf
}

// Reference returns the value of `$ref` field
func (v *items) Reference() string {
	return v.reference
}

func (v *items) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *items) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *items) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.size = len(items)
	iter.list.items = items
	return &iter
}

func (v *items) Validate(recurse bool) error {
	return newValidator(recurse).Validate(context.Background(), v)
}
