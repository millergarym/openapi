package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T16:03:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *SecurityRequirementList) Clear() error {
	*v = SecurityRequirementList(nil)
	return nil
}

func (v SecurityRequirementList) Resolve(resolver *Resolver) error {
	if len(v) > 0 {
		for i, elem := range v {
			if err := elem.Resolve(resolver); err != nil {
				return errors.Wrapf(err, `failed to resolve SecurityRequirementList (index = %d)`, i)
			}
		}
	}
	return nil
}

func (v *SecurityRequirementList) UnmarshalJSON(data []byte) error {
	var proxy []*securityRequirement
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}

	if len(proxy) == 0 {
		*v = SecurityRequirementList(nil)
		return nil
	}

	tmp := make(SecurityRequirementList, len(proxy))
	for i, value := range proxy {
		tmp[i] = value
	}
	*v = tmp
	return nil
}
