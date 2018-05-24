package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T11:15:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// ServerVariableMutator is used to build an instance of ServerVariable. The user must
// call `Do()` after providing all the necessary information to
// the new instance of ServerVariable with new values
type ServerVariableMutator struct {
	proxy  *serverVariable
	target *serverVariable
}

// Do finalizes the matuation process for ServerVariable and returns the result
func (b *ServerVariableMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateServerVariable creates a new mutator object for ServerVariable
func MutateServerVariable(v ServerVariable) *ServerVariableMutator {
	return &ServerVariableMutator{
		target: v.(*serverVariable),
		proxy:  v.Clone().(*serverVariable),
	}
}

func (b *ServerVariableMutator) ClearEnum() *ServerVariableMutator {
	b.proxy.enum.Clear()
	return b
}

func (b *ServerVariableMutator) Enum(value string) *ServerVariableMutator {
	b.proxy.enum = append(b.proxy.enum, value)
	return b
}

// Default sets the Default field for object ServerVariable.
func (b *ServerVariableMutator) Default(v string) *ServerVariableMutator {
	b.proxy.defaultValue = v
	return b
}

// Description sets the Description field for object ServerVariable.
func (b *ServerVariableMutator) Description(v string) *ServerVariableMutator {
	b.proxy.description = v
	return b
}
