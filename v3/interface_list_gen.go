package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T13:02:46+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

func (v *InterfaceList) Clear() error {
	*v = InterfaceList(nil)
	return nil
}

func (v InterfaceList) Resolve(resolver *Resolver) error {
	return nil
}