package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T20:40:48+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sort"
)

var _ = sort.Strings

func (v *securityScheme) Type() string {
	return v.typ
}

func (v *securityScheme) Description() string {
	return v.description
}

func (v *securityScheme) Name() string {
	return v.name
}

func (v *securityScheme) In() string {
	return v.in
}

func (v *securityScheme) Flow() string {
	return v.flow
}

func (v *securityScheme) AuthorizationURL() string {
	return v.authorizationURL
}

func (v *securityScheme) TokenURL() string {
	return v.tokenURL
}

func (v *securityScheme) Scopes() *StringMapIterator {
	var keys []string
	for key := range v.scopes {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.scopes[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter StringMapIterator
	iter.list.items = items
	return &iter
}

func (v *securityScheme) Reference() string {
	return v.reference
}

func (v *securityScheme) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *securityScheme) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *securityScheme) Validate() error {
	return nil
}