package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T20:40:48+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// HeaderBuilder is used to build an instance of Header. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Header
type HeaderBuilder struct {
	target *header
}

// Do finalizes the building process for Header and returns the result
func (b *HeaderBuilder) Do() (Header, error) {
	if err := b.target.Validate(); err != nil {
		return nil, errors.Wrap(err, `validation failed`)
	}
	return b.target, nil
}

// NewHeader creates a new builder object for Header
func NewHeader(typ string) *HeaderBuilder {
	return &HeaderBuilder{
		target: &header{
			typ: typ,
		},
	}
}

// Name sets the Name field for object Header.
func (b *HeaderBuilder) Name(v string) *HeaderBuilder {
	b.target.name = v
	return b
}

// Description sets the Description field for object Header.
func (b *HeaderBuilder) Description(v string) *HeaderBuilder {
	b.target.description = v
	return b
}

// Format sets the Format field for object Header.
func (b *HeaderBuilder) Format(v string) *HeaderBuilder {
	b.target.format = v
	return b
}

// Items sets the Items field for object Header.
func (b *HeaderBuilder) Items(v Items) *HeaderBuilder {
	b.target.items = v
	return b
}

// CollectionFormat sets the CollectionFormat field for object Header.
func (b *HeaderBuilder) CollectionFormat(v CollectionFormat) *HeaderBuilder {
	b.target.collectionFormat = v
	return b
}

// DefaultValue sets the DefaultValue field for object Header.
func (b *HeaderBuilder) DefaultValue(v interface{}) *HeaderBuilder {
	b.target.defaultValue = v
	return b
}

// Maximum sets the Maximum field for object Header.
func (b *HeaderBuilder) Maximum(v float64) *HeaderBuilder {
	b.target.maximum = v
	return b
}

// ExclusiveMaximum sets the ExclusiveMaximum field for object Header.
func (b *HeaderBuilder) ExclusiveMaximum(v float64) *HeaderBuilder {
	b.target.exclusiveMaximum = v
	return b
}

// Minimum sets the Minimum field for object Header.
func (b *HeaderBuilder) Minimum(v float64) *HeaderBuilder {
	b.target.minimum = v
	return b
}

// ExclusiveMinimum sets the ExclusiveMinimum field for object Header.
func (b *HeaderBuilder) ExclusiveMinimum(v float64) *HeaderBuilder {
	b.target.exclusiveMinimum = v
	return b
}

// MaxLength sets the MaxLength field for object Header.
func (b *HeaderBuilder) MaxLength(v int) *HeaderBuilder {
	b.target.maxLength = v
	return b
}

// MinLength sets the MinLength field for object Header.
func (b *HeaderBuilder) MinLength(v int) *HeaderBuilder {
	b.target.minLength = v
	return b
}

// Pattern sets the Pattern field for object Header.
func (b *HeaderBuilder) Pattern(v string) *HeaderBuilder {
	b.target.pattern = v
	return b
}

// MaxItems sets the MaxItems field for object Header.
func (b *HeaderBuilder) MaxItems(v int) *HeaderBuilder {
	b.target.maxItems = v
	return b
}

// MinItems sets the MinItems field for object Header.
func (b *HeaderBuilder) MinItems(v int) *HeaderBuilder {
	b.target.minItems = v
	return b
}

// UniqueItems sets the UniqueItems field for object Header.
func (b *HeaderBuilder) UniqueItems(v bool) *HeaderBuilder {
	b.target.uniqueItems = v
	return b
}

// Enum sets the Enum field for object Header.
func (b *HeaderBuilder) Enum(v ...interface{}) *HeaderBuilder {
	b.target.enum = v
	return b
}

// MultipleOf sets the MultipleOf field for object Header.
func (b *HeaderBuilder) MultipleOf(v float64) *HeaderBuilder {
	b.target.multipleOf = v
	return b
}

// Reference sets the $ref (reference) field for object Header.
func (b *HeaderBuilder) Reference(v string) *HeaderBuilder {
	b.target.reference = v
	return b
}