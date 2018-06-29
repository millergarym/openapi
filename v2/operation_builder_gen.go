package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// OperationBuilder is used to build an instance of Operation. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Operation
type OperationBuilder struct {
	target *operation
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *OperationBuilder) MustBuild(options ...Option) Operation {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Operation and returns the result
func (b *OperationBuilder) Build(options ...Option) (Operation, error) {
	validate := true
	for _, option := range options {
		switch option.Name() {
		case optkeyValidate:
			validate = option.Value().(bool)
		}
	}
	if validate {
		if err := b.target.Validate(false); err != nil {
			return nil, errors.Wrap(err, `validation failed`)
		}
	}
	return b.target, nil
}

// NewOperation creates a new builder object for Operation
func NewOperation(responses Responses) *OperationBuilder {
	return &OperationBuilder{
		target: &operation{
			responses: responses,
		},
	}
}

// Tags sets the tags field for object Operation.
func (b *OperationBuilder) Tags(v ...string) *OperationBuilder {
	b.target.tags = v
	return b
}

// Summary sets the summary field for object Operation.
func (b *OperationBuilder) Summary(v string) *OperationBuilder {
	b.target.summary = v
	return b
}

// Description sets the description field for object Operation.
func (b *OperationBuilder) Description(v string) *OperationBuilder {
	b.target.description = v
	return b
}

// ExternalDocs sets the externalDocs field for object Operation.
func (b *OperationBuilder) ExternalDocs(v ExternalDocumentation) *OperationBuilder {
	b.target.externalDocs = v
	return b
}

// OperationID sets the operationID field for object Operation.
func (b *OperationBuilder) OperationID(v string) *OperationBuilder {
	b.target.operationID = v
	return b
}

// Consumes sets the consumes field for object Operation.
func (b *OperationBuilder) Consumes(v ...MIMEType) *OperationBuilder {
	b.target.consumes = v
	return b
}

// Produces sets the produces field for object Operation.
func (b *OperationBuilder) Produces(v ...MIMEType) *OperationBuilder {
	b.target.produces = v
	return b
}

// Parameters sets the parameters field for object Operation.
func (b *OperationBuilder) Parameters(v ...Parameter) *OperationBuilder {
	b.target.parameters = v
	return b
}

// Schemes sets the schemes field for object Operation.
func (b *OperationBuilder) Schemes(v ...Scheme) *OperationBuilder {
	b.target.schemes = v
	return b
}

// Deprecated sets the deprecated field for object Operation.
func (b *OperationBuilder) Deprecated(v bool) *OperationBuilder {
	b.target.deprecated = v
	return b
}

// Security sets the security field for object Operation.
func (b *OperationBuilder) Security(v ...SecurityRequirement) *OperationBuilder {
	b.target.security = v
	return b
}

// Reference sets the $ref (reference) field for object Operation.
func (b *OperationBuilder) Reference(v string) *OperationBuilder {
	b.target.reference = v
	return b
}

// Extension sets an arbitrary element (an extension) to the
// object Operation. The extension name should start with a "x-"
func (b *OperationBuilder) Extension(name string, value interface{}) *OperationBuilder {
	b.target.extensions[name] = value
	return b
}
