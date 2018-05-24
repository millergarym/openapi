package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T11:15:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"log"
)

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"
)

var _ = errors.Cause

type callbackMarshalProxy struct {
	Reference string              `json:"$ref,omitempty"`
	URLs      map[string]PathItem `json:""`
}

type callbackUnmarshalProxy struct {
	Reference string                     `json:"$ref,omitempty"`
	URLs      map[string]json.RawMessage `json:""`
}

func (v *callback) MarshalJSON() ([]byte, error) {
	var proxy callbackMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.URLs = v.uRLs
	return json.Marshal(proxy)
}

func (v *callback) UnmarshalJSON(data []byte) error {
	var proxy callbackUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}

	if len(proxy.URLs) > 0 {
		m := make(map[string]PathItem)
		for key, pv := range proxy.URLs {
			var decoded pathItem
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field URLs`, key)
			}
			m[key] = &decoded
		}
		v.uRLs = m
	}
	return nil
}

func (v *callback) Resolve(resolver *Resolver) error {
	log.Printf(`callback.Resolve`)
	if v.IsReference() {
		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*callback)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Callback, but got %T`, resolved)
		}
		mutator := MutateCallback(v)
		mutator.URLs(asserted.URLs())
		return errors.Wrap(mutator.Do(), `failed to mutate`)
	}
	return nil
}

func (v *callback) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "uRLs":
		target = v.uRLs
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
