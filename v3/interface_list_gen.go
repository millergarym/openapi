package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T07:35:18+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *InterfaceList) Clear() error {
	*v = InterfaceList(nil)
	return nil
}

func (v InterfaceList) Resolve(resolver *Resolver) error {
	return nil
}
