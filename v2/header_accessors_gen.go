package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T21:36:31+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *header) Name() string {
	return v.name
}

func (v *header) Description() string {
	return v.description
}

func (v *header) Type() string {
	return v.typ
}

func (v *header) Format() string {
	return v.format
}

func (v *header) Items() Items {
	return v.items
}

func (v *header) CollectionFormat() CollectionFormat {
	return v.collectionFormat
}

func (v *header) DefaultValue() interface{} {
	return v.defaultValue
}

func (v *header) Maximum() float64 {
	return v.maximum
}

func (v *header) ExclusiveMaximum() float64 {
	return v.exclusiveMaximum
}

func (v *header) Minimum() float64 {
	return v.minimum
}

func (v *header) ExclusiveMinimum() float64 {
	return v.exclusiveMinimum
}

func (v *header) MaxLength() int {
	return v.maxLength
}

func (v *header) MinLength() int {
	return v.minLength
}

func (v *header) Pattern() string {
	return v.pattern
}

func (v *header) MaxItems() int {
	return v.maxItems
}

func (v *header) MinItems() int {
	return v.minItems
}

func (v *header) UniqueItems() bool {
	return v.uniqueItems
}

func (v *header) Enum() *InterfaceListIterator {
	var items []interface{}
	for _, item := range v.enum {
		items = append(items, item)
	}
	var iter InterfaceListIterator
	iter.items = items
	return &iter
}

func (v *header) MultipleOf() float64 {
	return v.multipleOf
}

func (v *header) Reference() string {
	return v.reference
}

func (v *header) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *header) Validate() error {
	return nil
}
