package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T11:15:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

func (v *StringList) Clear() error {
	*v = StringList(nil)
	return nil
}

func (v StringList) Resolve(resolver *Resolver) error {
	return nil
}
