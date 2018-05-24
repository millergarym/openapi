package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T11:15:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// ServerBuilder is used to build an instance of Server. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Server
type ServerBuilder struct {
	target *server
}

// Do finalizes the building process for Server and returns the result
func (b *ServerBuilder) Do() Server {
	return b.target
}

// NewServer creates a new builder object for Server
func NewServer(uRL string) *ServerBuilder {
	return &ServerBuilder{
		target: &server{
			uRL: uRL,
		},
	}
}

// Description sets the Description field for object Server.
func (b *ServerBuilder) Description(v string) *ServerBuilder {
	b.target.description = v
	return b
}

// Variables sets the Variables field for object Server.
func (b *ServerBuilder) Variables(v map[string]ServerVariable) *ServerBuilder {
	b.target.variables = v
	return b
}

// Reference sets the $ref (reference) field for object Server.
func (b *ServerBuilder) Reference(v string) *ServerBuilder {
	b.target.reference = v
	return b
}
