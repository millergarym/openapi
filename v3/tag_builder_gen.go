package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:44:15+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// TagBuilder is used to build an instance of Tag. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Tag
type TagBuilder struct {
	target *tag
}

// Build finalizes the building process for Tag and returns the result
func (b *TagBuilder) Build() Tag {
	return b.target
}

// NewTag creates a new builder object for Tag
func NewTag(name string) *TagBuilder {
	return &TagBuilder{
		target: &tag{
			name: name,
		},
	}
}

// Description sets the Description field for object Tag.
func (b *TagBuilder) Description(v string) *TagBuilder {
	b.target.description = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object Tag.
func (b *TagBuilder) ExternalDocs(v ExternalDocumentation) *TagBuilder {
	b.target.externalDocs = v
	return b
}
