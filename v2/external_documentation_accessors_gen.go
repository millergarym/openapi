package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T21:36:31+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *externalDocumentation) URL() string {
	return v.uRL
}

func (v *externalDocumentation) Description() string {
	return v.description
}

func (v *externalDocumentation) Reference() string {
	return v.reference
}

func (v *externalDocumentation) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *externalDocumentation) Validate() error {
	return nil
}
