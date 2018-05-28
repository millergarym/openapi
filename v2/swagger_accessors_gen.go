package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T16:03:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

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
	var items []interface{}
	for key, item := range v.definitions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter SchemaMapIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) Parameters() *ParameterMapIterator {
	var items []interface{}
	for key, item := range v.parameters {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ParameterMapIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) Responses() *ResponseMapIterator {
	var items []interface{}
	for key, item := range v.responses {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ResponseMapIterator
	iter.list.items = items
	return &iter
}

func (v *swagger) SecurityDefinitions() *SecuritySchemeMapIterator {
	var items []interface{}
	for key, item := range v.securityDefinitions {
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

func (v *swagger) Validate() error {
	return nil
}
