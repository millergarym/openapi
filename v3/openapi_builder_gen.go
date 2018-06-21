package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// OpenAPIBuilder is used to build an instance of OpenAPI. The user must
// call `Do()` after providing all the necessary information to
// build an instance of OpenAPI
type OpenAPIBuilder struct {
	target *openAPI
}

// Do finalizes the building process for OpenAPI and returns the result
func (b *OpenAPIBuilder) Do() OpenAPI {
	return b.target
}

// NewOpenAPI creates a new builder object for OpenAPI
func NewOpenAPI(info Info, paths Paths) *OpenAPIBuilder {
	return &OpenAPIBuilder{
		target: &openAPI{
			version: DefaultVersion,
			info:    info,
			paths:   paths,
		},
	}
}

// Version sets the Version field for object OpenAPI. If this is not called,
// a default value (DefaultVersion) is assigned to this field
func (b *OpenAPIBuilder) Version(v string) *OpenAPIBuilder {
	b.target.version = v
	return b
}

// Servers sets the Servers field for object OpenAPI.
func (b *OpenAPIBuilder) Servers(v []Server) *OpenAPIBuilder {
	b.target.servers = v
	return b
}

// Components sets the Components field for object OpenAPI.
func (b *OpenAPIBuilder) Components(v Components) *OpenAPIBuilder {
	b.target.components = v
	return b
}

// Security sets the Security field for object OpenAPI.
func (b *OpenAPIBuilder) Security(v SecurityRequirement) *OpenAPIBuilder {
	b.target.security = v
	return b
}

// Tags sets the Tags field for object OpenAPI.
func (b *OpenAPIBuilder) Tags(v []Tag) *OpenAPIBuilder {
	b.target.tags = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object OpenAPI.
func (b *OpenAPIBuilder) ExternalDocs(v ExternalDocumentation) *OpenAPIBuilder {
	b.target.externalDocs = v
	return b
}

// Reference sets the $ref (reference) field for object OpenAPI.
func (b *OpenAPIBuilder) Reference(v string) *OpenAPIBuilder {
	b.target.reference = v
	return b
}
