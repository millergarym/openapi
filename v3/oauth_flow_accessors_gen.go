package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T21:43:39+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *oAuthFlow) AuthorizationURL() string {
	return v.authorizationURL
}

func (v *oAuthFlow) TokenURL() string {
	return v.tokenURL
}

func (v *oAuthFlow) RefreshURL() string {
	return v.refreshURL
}

func (v *oAuthFlow) Scopes() map[string]string {
	return v.scopes
}
