package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T21:43:39+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type tagMarshalProxy struct {
	Name         string                `json:"name" builder:"required"`
	Description  string                `json:"description,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type tagUnmarshalProxy struct {
	Name         string          `json:"name" builder:"required"`
	Description  string          `json:"description,omitempty"`
	ExternalDocs json.RawMessage `json:"externalDocs,omitempty"`
}

func (v *tag) MarshalJSON() ([]byte, error) {
	var proxy tagMarshalProxy
	proxy.Name = v.name
	proxy.Description = v.description
	proxy.ExternalDocs = v.externalDocs
	return json.Marshal(proxy)
}

func (v *tag) UnmarshalJSON(data []byte) error {
	var proxy tagUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.name = proxy.Name
	v.description = proxy.Description

	if len(proxy.ExternalDocs) > 0 {
		var decoded externalDocumentation
		if err := json.Unmarshal(proxy.ExternalDocs, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		v.externalDocs = &decoded
	}
	return nil
}
