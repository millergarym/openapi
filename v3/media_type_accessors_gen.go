package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T07:35:18+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *mediaType) Name() string {
	return v.name
}

func (v *mediaType) Mime() string {
	return v.mime
}

func (v *mediaType) Schema() Schema {
	return v.schema
}

func (v *mediaType) Examples() *ExampleMapIterator {
	var items []interface{}
	for key, item := range v.examples {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExampleMapIterator
	iter.list.items = items
	return &iter
}

func (v *mediaType) Encoding() *EncodingMapIterator {
	var items []interface{}
	for key, item := range v.encoding {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter EncodingMapIterator
	iter.list.items = items
	return &iter
}

func (v *mediaType) Reference() string {
	return v.reference
}

func (v *mediaType) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
