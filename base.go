// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package keyboard

import (
	"strings"
	"time"
)

// Binding represents a single key/sequence binding.
type Binding struct {
	str     string // String version of the binding.
	Keys    []Key  // Keys involved in the binding. This includes modifiers.
	Handler func() // Handler function to be called when binding is triggered.
}

func (b *Binding) String() string { return b.str }

// Base is the base type for all backend implementations.
// It takes care of some common house keeping and ensures they all
// qualify as a Keyboard interface.
type Base struct {
	bindings []*Binding    // Set of registered bindings.
	record   []Key         // List of recorded keypresses.
	stamp    time.Time     // Last key press.
	timeout  time.Duration // Timeout for sequence resets.
}

// NewBase creates a new Base instance.
func NewBase() *Base {
	b := new(Base)
	b.timeout = time.Second >> 1
	b.Clear()
	return b
}

// Bindings returns a list of current bindings.
func (b *Base) Bindings() []*Binding { return b.bindings }

// SetTimeout sets the timeout for sequence resets in nanoseconds
func (b *Base) SetTimeout(d int64) { b.timeout = time.Duration(d) }

// RecordKey records the given key press.
func (b *Base) RecordKey(key Key, mods Modifier) {
	if time.Since(b.stamp) > b.timeout {
		b.record = b.record[:0]
	}

	b.stamp = time.Now()
	b.record = append(b.record, Key(mods)<<8|key)

	// If we have a binding match, invoke its handler.
	binding := b.fullMatch()
	if binding != nil {
		b.record = b.record[:0]
		binding.Handler()
		return
	}

	// If no match is possible with the current recording, clear the buffer.
	if !b.partialMatch() {
		b.record = b.record[:0]
		return
	}
}

// partialMatch returns true if the currently recorded sequence
// is a partial match for one of our bindings.
func (b *Base) partialMatch() bool {
	rlen := len(b.record)

	for _, bind := range b.bindings {
		if rlen < len(bind.Keys) && matchList(b.record, bind.Keys[:rlen]) {
			return true
		}
	}

	return false
}

// fullMatch checks of the currently recorded sequence matches one of our bindings.
func (b *Base) fullMatch() *Binding {
	rlen := len(b.record)

	for _, bind := range b.bindings {
		if rlen == len(bind.Keys) && matchList(b.record, bind.Keys) {
			return bind
		}
	}

	return nil
}

// Bind binds the given keys or sequences to the specified handler.
func (b *Base) Bind(handler func(), keys ...string) {
	if len(keys) == 0 {
		return
	}

	for i := range keys {
		b.bindings = append(b.bindings, &Binding{
			str:     keys[i],
			Keys:    parseKeys(keys[i]),
			Handler: handler,
		})
	}
}

// Unbind removes the binding for the given key or sequence.
func (b *Base) Unbind(key string) {
	idx := b.index(key)
	if idx == -1 {
		return
	}

	b.bindings[idx].Handler = nil
	b.bindings[idx].Keys = nil

	copy(b.bindings[idx:], b.bindings[idx+1:])
	b.bindings = b.bindings[:len(b.bindings)-1]
}

// Clear removes all bindings.
func (b *Base) Clear() {
	for i, bind := range b.bindings {
		bind.Handler = nil
		bind.Keys = nil
		b.bindings[i] = nil
	}

	b.bindings = b.bindings[:0]
}

// Call invokes the handler for the given key or sequence.
func (b *Base) Call(key string) {
	idx := b.index(key)

	if idx != -1 {
		b.bindings[idx].Handler()
	}
}

// Poll can be called with an backend-specific event object
// in cases where this is necessary. The backend should process the
// event accordingly. Not all backends will use this.
func (*Base) Poll(interface{}) {}

// index returns the index of the given binding.
func (b *Base) index(key string) int {
	for i, b := range b.bindings {
		if strings.EqualFold(b.str, key) {
			return i
		}
	}

	return -1
}

// matchList returns true if the two list contents match. 
func matchList(a, b []Key) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
