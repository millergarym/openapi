package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sync"
)

// PathItemMutator is used to build an instance of PathItem. The user must
// call `Apply()` after providing all the necessary information to
// the new instance of PathItem with new values
type PathItemMutator struct {
	lock   sync.Locker
	proxy  *pathItem
	target *pathItem
}

// Apply finalizes the matuation process for PathItem and returns the result
func (m *PathItemMutator) Apply() error {
	m.lock.Lock()
	defer m.lock.Unlock()
	*m.target = *m.proxy
	return nil
}

// MutatePathItem creates a new mutator object for PathItem
// Operations on the mutator are safe to be used concurrently, except for
// when calling `Apply()`, where the user is responsible for restricting access
// to the target object to be mutated
func MutatePathItem(v PathItem, options ...Option) *PathItemMutator {
	var lock sync.Locker = &sync.Mutex{}
	for _, option := range options {
		switch option.Name() {
		case optkeyLocker:
			lock = option.Value().(sync.Locker)
		}
	}
	if lock == nil {
		lock = nilLock{}
	}
	return &PathItemMutator{
		lock:   lock,
		target: v.(*pathItem),
		proxy:  v.Clone().(*pathItem),
	}
}

// ClearParameters clears all elements in parameters
func (m *PathItemMutator) ClearParameters() *PathItemMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	_ = m.proxy.parameters.Clear()
	return m
}

// Parameter appends a value to parameters
func (m *PathItemMutator) Parameter(value Parameter) *PathItemMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.parameters = append(m.proxy.parameters, value)
	return m
}

// Extension sets an arbitrary extension field in PathItem
func (m *PathItemMutator) Extension(name string, value interface{}) *PathItemMutator {
	if m.proxy.extensions == nil {
		m.proxy.extensions = Extensions{}
	}
	m.proxy.extensions[name] = value
	return m
}
