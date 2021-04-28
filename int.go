package state

import (
	"github.com/bytedisciple/logger"
	"unsafe"
)

// Int - Exposed Interface of int type
type Int interface {
	Get() int
	Name() string
	Set(int)
	Sub(onChange *func(oldValue, newValue int))
	Unsub(onChange *func(oldValue, newValue int))
}

// intState private internal int representation
type intState struct {
	i    int
	name string
	subs map[uintptr]func(oldValue, newValue int)
}

func NewInt(name string) Int {
	return intState{
		i:    0,
		name: name,
		subs: map[uintptr]func(oldValue int, newValue int){},
	}
}

func (i intState) Get() int {
	return i.i
}

func (i intState) Name() string {
	return i.name
}

func (i intState) Set(newValue int) {
	oldValue := i.i
	for key, val := range i.subs {
		logger.Debugf("Running update function on callback [%v] for object %v", key, i.name)
		val(oldValue, newValue)
	}
	i.i = newValue
}

// Sub - Subscribe to changes on this object by passing your callback function to the object.
// This function uses the memory address of the function as a key for the internal map.
// For this reason it accepts a pointer to the function as to not allow anonymous functions
// which would all have the same memory address.
func (i intState) Sub(onChange *func(oldValue, newValue int)) {
	key := uintptr(unsafe.Pointer(onChange))

	_, exists := i.subs[key]
	if exists {
		logger.Debugf("Key %v exists for object %s, overriding with new function", key, i.name)
	}

	i.subs[key] = *onChange
}

func (i intState) Unsub(onChange *func(oldValue, newValue int)) {
	key := uintptr(unsafe.Pointer(onChange))

	_, exists := i.subs[key]
	if exists {
		logger.Debugf("Key %v exists for object %s, deleting", key, i.name)
	}

	delete(i.subs, key)
}
