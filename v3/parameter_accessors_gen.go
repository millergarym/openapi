package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:44:15+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *parameter) Name() string {
	return v.name
}

func (v *parameter) In() Location {
	return v.in
}

func (v *parameter) Required() bool {
	return v.required
}

func (v *parameter) Description() string {
	return v.description
}

func (v *parameter) Deprecated() bool {
	return v.deprecated
}

func (v *parameter) AllowEmptyValue() bool {
	return v.allowEmptyValue
}

func (v *parameter) Explode() bool {
	return v.explode
}

func (v *parameter) AllowReserved() bool {
	return v.allowReserved
}

func (v *parameter) Schema() Schema {
	return v.schema
}

func (v *parameter) Example() interface{} {
	return v.example
}

func (v *parameter) Examples() map[string]Example {
	return v.examples
}

func (v *parameter) Content() map[string]MediaType {
	return v.content
}
