package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sort"
)

var _ = sort.Strings

func (v *swagger) Version() string {
	return v.version
}

func (v *swagger) Info() Info {
	return v.info
}

func (v *swagger) Host() string {
	return v.host
}

func (v *swagger) BasePath() string {
	return v.basePath
}

func (v *swagger) Schemes() *SchemeListIterator {
	var items []interface{}
	for _, item := range v.schemes {
		items = append(items, item)
	}
	var iter SchemeListIterator
	iter.items = items
	return &iter
}

func (v *swagger) Consumes() *MIMETypeListIterator {
	var items []interface{}
	for _, item := range v.consumes {
		items = append(items, item)
	}
	var iter MIMETypeListIterator
	iter.items = items
	return &iter
}

func (v *swagger) Produces() *MIMETypeListIterator {
	var items []interface{}
	for _, item := range v.produces {
		items = append(items, item)
	}
	var iter MIMETypeListIterator
	iter.items = items
	return &iter
}

func (v *swagger) Paths() Paths {
	return v.paths
}

func (v *swagger) Definitions() *SchemaMapIterator {
	var keys []string
	for key := range v.definitions {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.definitions[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter SchemaMapIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) Parameters() *ParameterMapIterator {
	var keys []string
	for key := range v.parameters {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.parameters[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ParameterMapIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) Responses() *ResponseMapIterator {
	var keys []string
	for key := range v.responses {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.responses[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ResponseMapIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) SecurityDefinitions() *SecuritySchemeMapIterator {
	var keys []string
	for key := range v.securityDefinitions {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.securityDefinitions[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter SecuritySchemeMapIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) Security() *SecurityRequirementListIterator {
	var items []interface{}
	for _, item := range v.security {
		items = append(items, item)
	}
	var iter SecurityRequirementListIterator
	iter.items = items
	return &iter
}

func (v *swagger) Tags() *TagListIterator {
	var items []interface{}
	for _, item := range v.tags {
		items = append(items, item)
	}
	var iter TagListIterator
	iter.items = items
	return &iter
}

func (v *swagger) ExternalDocs() ExternalDocumentation {
	return v.externalDocs
}

func (v *swagger) Reference() string {
	return v.reference
}

func (v *swagger) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *swagger) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

func (v *swagger) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}
