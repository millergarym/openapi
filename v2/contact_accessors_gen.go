package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T16:03:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *contact) Name() string {
	return v.name
}

func (v *contact) URL() string {
	return v.uRL
}

func (v *contact) Email() string {
	return v.email
}

func (v *contact) Reference() string {
	return v.reference
}

func (v *contact) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *contact) Validate() error {
	return nil
}
