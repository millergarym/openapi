package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T21:43:39+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// RequestBodyBuilder is used to build an instance of RequestBody. The user must
// call `Do()` after providing all the necessary information to
// build an instance of RequestBody
type RequestBodyBuilder struct {
	target *requestBody
}

// Build finalizes the building process for RequestBody and returns the result
func (b *RequestBodyBuilder) Do() RequestBody {
	return b.target
}

// NewRequestBody creates a new builder object for RequestBody
func NewRequestBody() *RequestBodyBuilder {
	return &RequestBodyBuilder{
		target: &requestBody{},
	}
}

// Description sets the Description field for object RequestBody.
func (b *RequestBodyBuilder) Description(v string) *RequestBodyBuilder {
	b.target.description = v
	return b
}

// Content sets the Content field for object RequestBody.
func (b *RequestBodyBuilder) Content(v map[string]MediaType) *RequestBodyBuilder {
	b.target.content = v
	return b
}

// Required sets the Required field for object RequestBody.
func (b *RequestBodyBuilder) Required(v bool) *RequestBodyBuilder {
	b.target.required = v
	return b
}
