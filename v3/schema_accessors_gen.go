package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:44:15+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *schema) Reference() string {
	return v.reference
}

func (v *schema) Title() string {
	return v.title
}

func (v *schema) MultipleOf() float64 {
	return v.multipleOf
}

func (v *schema) Maximum() float64 {
	return v.maximum
}

func (v *schema) ExclusiveMaximum() float64 {
	return v.exclusiveMaximum
}

func (v *schema) Minimum() float64 {
	return v.minimum
}

func (v *schema) ExclusiveMinimum() float64 {
	return v.exclusiveMinimum
}

func (v *schema) MaxLength() int {
	return v.maxLength
}

func (v *schema) MinLength() int {
	return v.minLength
}

func (v *schema) Pattern() string {
	return v.pattern
}

func (v *schema) MaxItems() int {
	return v.maxItems
}

func (v *schema) MinItems() int {
	return v.minItems
}

func (v *schema) UniqueItems() bool {
	return v.uniqueItems
}

func (v *schema) MaxProperties() int {
	return v.maxProperties
}

func (v *schema) MinProperties() int {
	return v.minProperties
}

func (v *schema) Required() []string {
	return v.required
}

func (v *schema) Enum() []interface{} {
	return v.enum
}

func (v *schema) Type() PrimitiveType {
	return v.typ
}

func (v *schema) AllOf() []Schema {
	return v.allOf
}

func (v *schema) OneOf() []Schema {
	return v.oneOf
}

func (v *schema) AnyOf() []Schema {
	return v.anyOf
}

func (v *schema) Not() Schema {
	return v.not
}

func (v *schema) Items() Schema {
	return v.items
}

func (v *schema) Properties() map[string]Schema {
	return v.properties
}

func (v *schema) Format() string {
	return v.format
}

func (v *schema) Default() interface{} {
	return v.defaultValue
}

func (v *schema) Nullable() bool {
	return v.nullable
}

func (v *schema) Discriminator() Discriminator {
	return v.discriminator
}

func (v *schema) ReadOnly() bool {
	return v.readOnly
}

func (v *schema) WriteOnly() bool {
	return v.writeOnly
}

func (v *schema) ExternalDocs() ExternalDocumentation {
	return v.externalDocs
}

func (v *schema) Example() interface{} {
	return v.example
}

func (v *schema) Deprecated() bool {
	return v.deprecated
}
