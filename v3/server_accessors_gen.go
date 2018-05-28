package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T07:35:18+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *server) URL() string {
	return v.uRL
}

func (v *server) Description() string {
	return v.description
}

func (v *server) Variables() *ServerVariableMapIterator {
	var items []interface{}
	for key, item := range v.variables {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ServerVariableMapIterator
	iter.list.items = items
	return &iter
}

func (v *server) Reference() string {
	return v.reference
}

func (v *server) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
