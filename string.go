package state

import (
	"github.com/bytedisciple/logger"
	"unsafe"
)

// String - Exposed Interface of int type
type String interface {
	Get() string
	Name() string
	Set(string)
	Sub(onChange *func(oldValue, newValue string))
	Unsub(onChange *func(oldValue, newValue string))
}

// stringState private internal string representation
type stringState struct {
	i    string
	name string
	subs map[uintptr]func(oldValue, newValue string)
}

func NewString(name string) String {
	return stringState{
		i:    "",
		name: name,
		subs: map[uintptr]func(oldValue string, newValue string){},
	}
}

func (s stringState) Get() string {
	return s.i
}

func (s stringState) Name() string {
	return s.name
}

func (s stringState) Set(newValue string) {
	oldValue := s.i
	for key, val := range s.subs {
		logger.Debugf("Running update function on callback [%v] for object %v", key, s.name)
		val(oldValue, newValue)
	}
	s.i = newValue
}

// Sub - Subscribe to changes on this object by passing your callback function to the object.
// This function uses the memory address of the function as a key for the internal map.
// For this reason it accepts a pointer to the function as to not allow anonymous functions
// which would all have the same memory address.
func (s stringState) Sub(onChange *func(oldValue, newValue string)) {
	key := uintptr(unsafe.Pointer(onChange))

	_, exists := s.subs[key]
	if exists {
		logger.Debugf("Key %v exists for object %s, overriding with new function", key, s.name)
	}

	s.subs[key] = *onChange
}

func (s stringState) Unsub(onChange *func(oldValue, newValue string)) {
	key := uintptr(unsafe.Pointer(onChange))

	_, exists := s.subs[key]
	if exists {
		logger.Debugf("Key %v exists for object %s, deleting", key, s.name)
	}

	delete(s.subs, key)
}
