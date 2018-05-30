package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sort"
)

var _ = sort.Strings

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

func (v *items) DefaultValue() interface{} {
	return v.defaultValue
}

func (v *items) Maximum() float64 {
	return v.maximum
}

func (v *items) ExclusiveMaximum() float64 {
	return v.exclusiveMaximum
}

func (v *items) Minimum() float64 {
	return v.minimum
}

func (v *items) ExclusiveMinimum() float64 {
	return v.exclusiveMinimum
}

func (v *items) MaxLength() int {
	return v.maxLength
}

func (v *items) MinLength() int {
	return v.minLength
}

func (v *items) Pattern() string {
	return v.pattern
}

func (v *items) MaxItems() int {
	return v.maxItems
}

func (v *items) MinItems() int {
	return v.minItems
}

func (v *items) UniqueItems() bool {
	return v.uniqueItems
}

func (v *items) Enum() *InterfaceListIterator {
	var items []interface{}
	for _, item := range v.enum {
		items = append(items, item)
	}
	var iter InterfaceListIterator
	iter.items = items
	return &iter
}

func (v *items) MultipleOf() float64 {
	return v.multipleOf
}

func (v *items) Reference() string {
	return v.reference
}

func (v *items) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *items) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

func (v *items) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *items) Validate() error {
	return nil
}
