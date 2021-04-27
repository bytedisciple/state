package state

import (
	"github.com/bytedisciple/logger"
	"testing"
)

func TestNewIntState(t *testing.T) {
	is := NewInt("myTestIntState")

	if is.Get() != 0 {
		t.Error("Default value should be 0")
	}

	f1 := func(oldValue, newValue int) {
		logger.Debugf("New value is: [%v]", newValue)
	}

	is.Sub(&f1)

	is.Set(1)
}

func TestIntManyCallbacks(t *testing.T){
	intState := NewInt("myTestIntState")

	size := 5

	bools := make([]bool, size)
	funcs := make([]func(oldValue, newValue int), size)


	for i := range funcs {
		logger.Infof("Creating function at position %v", i)
		// Must be reallocated in order to keep i from incrementing in the func below
		newIndex := i
		funcs[newIndex] = func(oldValue, newValue int) {
			 logger.Debugf("Running on position %v - %v", newIndex, &newIndex)
			 bools[newIndex] = true
		}

		intState.Sub(&funcs[i])
	}

	intState.Set(1)

	for i, v := range bools {
		if !v{
			t.Errorf("Position %v was not set to true!", i)
		}
	}
}