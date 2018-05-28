package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T21:36:31+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *parameter) Name() string {
	return v.name
}

func (v *parameter) Description() string {
	return v.description
}

func (v *parameter) Required() bool {
	return v.required
}

func (v *parameter) In() Location {
	return v.in
}

func (v *parameter) Schema() Schema {
	return v.schema
}

func (v *parameter) Type() PrimitiveType {
	return v.typ
}

func (v *parameter) Format() string {
	return v.format
}

func (v *parameter) Title() string {
	return v.title
}

func (v *parameter) AllowEmptyValue() bool {
	return v.allowEmptyValue
}

func (v *parameter) Items() Items {
	return v.items
}

func (v *parameter) CollectionFormat() CollectionFormat {
	return v.collectionFormat
}

func (v *parameter) DefaultValue() interface{} {
	return v.defaultValue
}

func (v *parameter) Maximum() float64 {
	return v.maximum
}

func (v *parameter) ExclusiveMaximum() float64 {
	return v.exclusiveMaximum
}

func (v *parameter) Minimum() float64 {
	return v.minimum
}

func (v *parameter) ExclusiveMinimum() float64 {
	return v.exclusiveMinimum
}

func (v *parameter) MaxLength() int {
	return v.maxLength
}

func (v *parameter) MinLength() int {
	return v.minLength
}

func (v *parameter) Pattern() string {
	return v.pattern
}

func (v *parameter) MaxItems() int {
	return v.maxItems
}

func (v *parameter) MinItems() int {
	return v.minItems
}

func (v *parameter) UniqueItems() bool {
	return v.uniqueItems
}

func (v *parameter) Enum() *InterfaceListIterator {
	var items []interface{}
	for _, item := range v.enum {
		items = append(items, item)
	}
	var iter InterfaceListIterator
	iter.items = items
	return &iter
}

func (v *parameter) MultipleOf() float64 {
	return v.multipleOf
}

func (v *parameter) Reference() string {
	return v.reference
}

func (v *parameter) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
