package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T07:35:18+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *example) Name() string {
	return v.name
}

func (v *example) Description() string {
	return v.description
}

func (v *example) Value() interface{} {
	return v.value
}

func (v *example) ExternalValue() string {
	return v.externalValue
}

func (v *example) Reference() string {
	return v.reference
}

func (v *example) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
