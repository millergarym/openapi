package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:44:15+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *server) URL() string {
	return v.uRL
}

func (v *server) Description() string {
	return v.description
}

func (v *server) Variables() map[string]ServerVariable {
	return v.variables
}
