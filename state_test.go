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

func TestName(t *testing.T) {
	name := "testname123"
	is := NewInt(name)
	if is.Name() != name {
		t.Errorf("Name not as expected!")
	}
}

func TestIntUnsub(t *testing.T){
	is := NewInt("int1")
	counter := 0

	f1 := func(oldValue, newValue int) {
		counter += newValue
	}

	is.Sub(&f1)
	is.Set(1)
	if counter != 1{
		t.Error("Inner callback function not run")
	}

	is.Unsub(&f1)
	is.Set(2)
	if counter != 1{
		t.Error("Counter should still be 1, callback should have been unregistered!")
	}

}

func TestOverrideFunc(t *testing.T) {
	is := NewInt("int1")
	counter := 0

	f1 := func(oldValue, newValue int) {
		counter += newValue
	}

	f1 = func(oldValue, newValue int) {
		counter -= newValue
	}

	is.Sub(&f1)
	is.Set(1)

	if counter != -1 {
		t.Error("f1 should have been overridden with its 2nd form that subtracts!")
	}

	// Has no effect
	is.Sub(&f1)
	is.Set(1)

	if counter != -2 {
		t.Error("f1 should have been overridden with its 2nd form that subtracts!")
	}
}