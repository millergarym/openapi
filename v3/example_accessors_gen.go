package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T13:02:46+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

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
