package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T16:03:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var _ = log.Printf
var _ = json.Unmarshal
var _ = errors.Cause

type responseMarshalProxy struct {
	Reference   string     `json:"$ref,omitempty"`
	Description string     `json:"description"`
	Schema      Schema     `json:"schema,omitempty"`
	Headers     HeaderMap  `json:"headers,omitempty"`
	Examples    ExampleMap `json:"example,omitempty"`
}

type responseUnmarshalProxy struct {
	Reference   string          `json:"$ref,omitempty"`
	Description string          `json:"description"`
	Schema      json.RawMessage `json:"schema,omitempty"`
	Headers     HeaderMap       `json:"headers,omitempty"`
	Examples    ExampleMap      `json:"example,omitempty"`
}

func (v *response) MarshalJSON() ([]byte, error) {
	var proxy responseMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Description = v.description
	proxy.Schema = v.schema
	proxy.Headers = v.headers
	proxy.Examples = v.examples
	return json.Marshal(proxy)
}

func (v *response) UnmarshalJSON(data []byte) error {
	var proxy responseUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.description = proxy.Description

	if len(proxy.Schema) > 0 {
		var decoded schema
		if err := json.Unmarshal(proxy.Schema, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Schema`)
		}

		v.schema = &decoded
	}
	v.headers = proxy.Headers
	v.examples = proxy.Examples
	return nil
}

func (v *response) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*response)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Response, but got %T`, resolved)
		}
		mutator := MutateResponse(v)
		mutator.Name(asserted.Name())
		mutator.Description(asserted.Description())
		mutator.Schema(asserted.Schema())
		for iter := asserted.Headers(); iter.Next(); {
			key, item := iter.Item()
			mutator.Header(key, item)
		}
		for iter := asserted.Examples(); iter.Next(); {
			key, item := iter.Item()
			mutator.Example(key, item)
		}
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.schema != nil {
		if err := v.schema.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Schema`)
		}
	}
	if v.headers != nil {
		if err := v.headers.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Headers`)
		}
	}
	if v.examples != nil {
		if err := v.examples.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Examples`)
		}
	}
	return nil
}

func (v *response) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "description":
		target = v.description
	case "schema":
		target = v.schema
	case "headers":
		target = v.headers
	case "example":
		target = v.examples
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
