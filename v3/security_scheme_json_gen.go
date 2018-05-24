package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T13:02:46+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var _ = log.Printf
var _ = errors.Cause

type securitySchemeMarshalProxy struct {
	Reference        string     `json:"$ref,omitempty"`
	Type             string     `json:"type"`
	Description      string     `json:"description"`
	Name             string     `json:"name"`
	In               string     `json:"in"`
	Scheme           string     `json:"scheme"`
	BearerFormat     string     `json:"bearerFormat"`
	Flows            OAuthFlows `json:"flows"`
	OpenIDConnectURL string     `json:"openIdConnectUrl"`
}

type securitySchemeUnmarshalProxy struct {
	Reference        string          `json:"$ref,omitempty"`
	Type             string          `json:"type"`
	Description      string          `json:"description"`
	Name             string          `json:"name"`
	In               string          `json:"in"`
	Scheme           string          `json:"scheme"`
	BearerFormat     string          `json:"bearerFormat"`
	Flows            json.RawMessage `json:"flows"`
	OpenIDConnectURL string          `json:"openIdConnectUrl"`
}

func (v *securityScheme) MarshalJSON() ([]byte, error) {
	var proxy securitySchemeMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Type = v.typ
	proxy.Description = v.description
	proxy.Name = v.name
	proxy.In = v.in
	proxy.Scheme = v.scheme
	proxy.BearerFormat = v.bearerFormat
	proxy.Flows = v.flows
	proxy.OpenIDConnectURL = v.openIDConnectURL
	return json.Marshal(proxy)
}

func (v *securityScheme) UnmarshalJSON(data []byte) error {
	var proxy securitySchemeUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.typ = proxy.Type
	v.description = proxy.Description
	v.name = proxy.Name
	v.in = proxy.In
	v.scheme = proxy.Scheme
	v.bearerFormat = proxy.BearerFormat

	if len(proxy.Flows) > 0 {
		var decoded oAuthFlows
		if err := json.Unmarshal(proxy.Flows, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Flows`)
		}

		v.flows = &decoded
	}
	v.openIDConnectURL = proxy.OpenIDConnectURL
	return nil
}

func (v *securityScheme) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*securityScheme)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type SecurityScheme, but got %T`, resolved)
		}
		mutator := MutateSecurityScheme(v)
		mutator.Type(asserted.Type())
		mutator.Description(asserted.Description())
		mutator.Name(asserted.Name())
		mutator.In(asserted.In())
		mutator.Scheme(asserted.Scheme())
		mutator.BearerFormat(asserted.BearerFormat())
		mutator.Flows(asserted.Flows())
		mutator.OpenIDConnectURL(asserted.OpenIDConnectURL())
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.flows != nil {
		if err := v.flows.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Flows`)
		}
	}
	return nil
}

func (v *securityScheme) QueryJSON(path string) (ret interface{}, ok bool) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	if path == "" {
		return v, true
	}

	var frag string
	if i := strings.Index(path, "/"); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ""
	}

	var target interface{}

	switch frag {
	case "type":
		target = v.typ
	case "description":
		target = v.description
	case "name":
		target = v.name
	case "in":
		target = v.in
	case "scheme":
		target = v.scheme
	case "bearerFormat":
		target = v.bearerFormat
	case "flows":
		target = v.flows
	case "openIdConnectUrl":
		target = v.openIDConnectURL
	default:
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}
	if path == "" {
		return target, true
	}
	return nil, false
}
