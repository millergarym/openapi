package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T16:03:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *operation) Tags() *TagListIterator {
	var items []interface{}
	for _, item := range v.tags {
		items = append(items, item)
	}
	var iter TagListIterator
	iter.items = items
	return &iter
}

func (v *operation) Summary() string {
	return v.summary
}

func (v *operation) Description() string {
	return v.description
}

func (v *operation) ExternalDocs() ExternalDocumentation {
	return v.externalDocs
}

func (v *operation) OperationID() string {
	return v.operationID
}

func (v *operation) Consumes() *MIMETypeListIterator {
	var items []interface{}
	for _, item := range v.consumes {
		items = append(items, item)
	}
	var iter MIMETypeListIterator
	iter.items = items
	return &iter
}

func (v *operation) Produces() *MIMETypeListIterator {
	var items []interface{}
	for _, item := range v.produces {
		items = append(items, item)
	}
	var iter MIMETypeListIterator
	iter.items = items
	return &iter
}

func (v *operation) Parameters() *ParameterListIterator {
	var items []interface{}
	for _, item := range v.parameters {
		items = append(items, item)
	}
	var iter ParameterListIterator
	iter.items = items
	return &iter
}

func (v *operation) Responses() Responses {
	return v.responses
}

func (v *operation) Schemes() *SchemeListIterator {
	var items []interface{}
	for _, item := range v.schemes {
		items = append(items, item)
	}
	var iter SchemeListIterator
	iter.items = items
	return &iter
}

func (v *operation) Deprecated() bool {
	return v.deprecated
}

func (v *operation) Security() *SecurityRequirementListIterator {
	var items []interface{}
	for _, item := range v.security {
		items = append(items, item)
	}
	var iter SecurityRequirementListIterator
	iter.items = items
	return &iter
}

func (v *operation) Reference() string {
	return v.reference
}

func (v *operation) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *operation) Validate() error {
	return nil
}
