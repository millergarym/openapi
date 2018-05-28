package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T21:36:31+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"log"
)

var _ = log.Printf

// OperationMutator is used to build an instance of Operation. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Operation with new values
type OperationMutator struct {
	proxy  *operation
	target *operation
}

// Do finalizes the matuation process for Operation and returns the result
func (b *OperationMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateOperation creates a new mutator object for Operation
func MutateOperation(v Operation) *OperationMutator {
	return &OperationMutator{
		target: v.(*operation),
		proxy:  v.Clone().(*operation),
	}
}

func (b *OperationMutator) ClearTags() *OperationMutator {
	b.proxy.tags.Clear()
	return b
}

func (b *OperationMutator) Tag(value Tag) *OperationMutator {
	b.proxy.tags = append(b.proxy.tags, value)
	return b
}

// Summary sets the Summary field for object Operation.
func (b *OperationMutator) Summary(v string) *OperationMutator {
	b.proxy.summary = v
	return b
}

// Description sets the Description field for object Operation.
func (b *OperationMutator) Description(v string) *OperationMutator {
	b.proxy.description = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object Operation.
func (b *OperationMutator) ExternalDocs(v ExternalDocumentation) *OperationMutator {
	b.proxy.externalDocs = v
	return b
}

// OperationID sets the OperationID field for object Operation.
func (b *OperationMutator) OperationID(v string) *OperationMutator {
	b.proxy.operationID = v
	return b
}

func (b *OperationMutator) ClearConsumes() *OperationMutator {
	b.proxy.consumes.Clear()
	return b
}

func (b *OperationMutator) Consume(value string) *OperationMutator {
	b.proxy.consumes = append(b.proxy.consumes, value)
	return b
}

func (b *OperationMutator) ClearProduces() *OperationMutator {
	b.proxy.produces.Clear()
	return b
}

func (b *OperationMutator) Produce(value string) *OperationMutator {
	b.proxy.produces = append(b.proxy.produces, value)
	return b
}

func (b *OperationMutator) ClearParameters() *OperationMutator {
	b.proxy.parameters.Clear()
	return b
}

func (b *OperationMutator) Parameter(value Parameter) *OperationMutator {
	b.proxy.parameters = append(b.proxy.parameters, value)
	return b
}

// Responses sets the Responses field for object Operation.
func (b *OperationMutator) Responses(v Responses) *OperationMutator {
	b.proxy.responses = v
	return b
}

func (b *OperationMutator) ClearSchemes() *OperationMutator {
	b.proxy.schemes.Clear()
	return b
}

func (b *OperationMutator) Scheme(value string) *OperationMutator {
	b.proxy.schemes = append(b.proxy.schemes, value)
	return b
}

// Deprecated sets the Deprecated field for object Operation.
func (b *OperationMutator) Deprecated(v bool) *OperationMutator {
	b.proxy.deprecated = v
	return b
}

func (b *OperationMutator) ClearSecurity() *OperationMutator {
	b.proxy.security.Clear()
	return b
}

func (b *OperationMutator) Security(value SecurityRequirement) *OperationMutator {
	b.proxy.security = append(b.proxy.security, value)
	return b
}
