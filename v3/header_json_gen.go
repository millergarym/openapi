package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T21:43:39+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type headerMarshalProxy struct {
	Required        bool                 `json:"required,omitempty"`
	Description     string               `json:"description,omitempty"`
	Deprecated      bool                 `json:"deprecated,omitempty"`
	AllowEmptyValue bool                 `json:"allowEmptyValue,omitempty"`
	Explode         bool                 `json:"explode,omitempty"`
	AllowReserved   bool                 `json:"allowReserved,omitempty"`
	Schema          Schema               `json:"schema,omitempty"`
	Example         interface{}          `json:"example,omitempty"`
	Examples        map[string]Example   `json:"examples,omitempty"`
	Content         map[string]MediaType `json:"content,omitempty"`
}

type headerUnmarshalProxy struct {
	Required        bool                       `json:"required,omitempty"`
	Description     string                     `json:"description,omitempty"`
	Deprecated      bool                       `json:"deprecated,omitempty"`
	AllowEmptyValue bool                       `json:"allowEmptyValue,omitempty"`
	Explode         bool                       `json:"explode,omitempty"`
	AllowReserved   bool                       `json:"allowReserved,omitempty"`
	Schema          json.RawMessage            `json:"schema,omitempty"`
	Example         interface{}                `json:"example,omitempty"`
	Examples        map[string]json.RawMessage `json:"examples,omitempty"`
	Content         map[string]json.RawMessage `json:"content,omitempty"`
}

func (v *header) MarshalJSON() ([]byte, error) {
	var proxy headerMarshalProxy
	proxy.Required = v.required
	proxy.Description = v.description
	proxy.Deprecated = v.deprecated
	proxy.AllowEmptyValue = v.allowEmptyValue
	proxy.Explode = v.explode
	proxy.AllowReserved = v.allowReserved
	proxy.Schema = v.schema
	proxy.Example = v.example
	proxy.Examples = v.examples
	proxy.Content = v.content
	return json.Marshal(proxy)
}

func (v *header) UnmarshalJSON(data []byte) error {
	var proxy headerUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.required = proxy.Required
	v.description = proxy.Description
	v.deprecated = proxy.Deprecated
	v.allowEmptyValue = proxy.AllowEmptyValue
	v.explode = proxy.Explode
	v.allowReserved = proxy.AllowReserved

	if len(proxy.Schema) > 0 {
		var decoded schema
		if err := json.Unmarshal(proxy.Schema, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Schema`)
		}

		v.schema = &decoded
	}
	v.example = proxy.Example

	if len(proxy.Examples) > 0 {
		m := make(map[string]Example)
		for key, pv := range proxy.Examples {
			var decoded example
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Examples`, key)
			}
			m[key] = &decoded
		}
		v.examples = m
	}

	if len(proxy.Content) > 0 {
		m := make(map[string]MediaType)
		for key, pv := range proxy.Content {
			var decoded mediaType
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Content`, key)
			}
			m[key] = &decoded
		}
		v.content = m
	}
	return nil
}
