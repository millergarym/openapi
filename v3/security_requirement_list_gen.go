package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T13:02:46+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

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
