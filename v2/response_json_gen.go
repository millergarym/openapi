package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
)

var _ = json.Unmarshal
var _ = fmt.Fprintf
var _ = log.Printf
var _ = strconv.ParseInt
var _ = errors.Cause

type responseMarshalProxy struct {
	Reference   string     `json:"$ref,omitempty"`
	Description string     `json:"description"`
	Schema      Schema     `json:"schema,omitempty"`
	Headers     HeaderMap  `json:"headers,omitempty"`
	Examples    ExampleMap `json:"example,omitempty"`
}

func (v *response) MarshalJSON() ([]byte, error) {
	var proxy responseMarshalProxy
	if s := v.reference; len(s) > 0 {
		return []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil
	}
	proxy.Description = v.description
	proxy.Schema = v.schema
	proxy.Headers = v.headers
	proxy.Examples = v.examples
	buf, err := json.Marshal(proxy)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal struct`)
	}
	if len(v.extensions) > 0 {
		extBuf, err := json.Marshal(v.extensions)
		if err != nil || len(extBuf) <= 2 {
			return nil, errors.Wrap(err, `failed to marshal struct (extensions)`)
		}
		buf = append(append(buf[:len(buf)-1], ','), extBuf[1:]...)
	}
	return buf, nil
}

func (v *response) UnmarshalJSON(data []byte) error {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if raw, ok := proxy["$ref"]; ok {
		if err := json.Unmarshal(raw, &v.reference); err != nil {
			return errors.Wrap(err, `failed to unmarshal $ref`)
		}
		return nil
	}

	mutator := MutateResponse(v)

	if raw, ok := proxy["description"]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field description`)
		}
		mutator.Description(decoded)
		delete(proxy, "description")
	}

	if raw, ok := proxy["schema"]; ok {
		var decoded schema
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Schema`)
		}

		mutator.Schema(&decoded)
		delete(proxy, "schema")
	}

	if raw, ok := proxy["headers"]; ok {
		var decoded HeaderMap
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Headers`)
		}
		for key, elem := range decoded {
			mutator.Header(key, elem)
		}
		delete(proxy, "headers")
	}

	if raw, ok := proxy["example"]; ok {
		var decoded ExampleMap
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Examples`)
		}
		for key, elem := range decoded {
			mutator.Example(key, elem)
		}
		delete(proxy, "example")
	}

	for name, raw := range proxy {
		if strings.HasPrefix(name, `x-`) {
			var ext interface{}
			if err := json.Unmarshal(raw, &ext); err != nil {
				return errors.Wrapf(err, `failed to unmarshal field %s`, name)
			}
			mutator.Extension(name, ext)
		}
	}

	if err := mutator.Do(); err != nil {
		return errors.Wrap(err, `failed to  unmarshal JSON`)
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

func ResponseFromJSON(buf []byte, v *Response) error {
	var tmp response
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}
	*v = &tmp
	return nil
}
