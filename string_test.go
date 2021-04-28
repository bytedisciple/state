package state

import (
	"github.com/bytedisciple/logger"
	"strconv"
	"testing"
)

func TestNewStringState(t *testing.T) {
	is := NewString("myTestStringState")

	if is.Get() != "" {
		t.Error("Default value should be \"\"")
	}

	f1 := func(oldValue, newValue string) {
		logger.Debugf("New value is: [%v]", newValue)
	}

	is.Sub(&f1)

	is.Set("a")
}

func TestStringManyCallbacks(t *testing.T){
	ss := NewString("myTestStringState")

	size := 5

	bools := make([]bool, size)
	funcs := make([]func(oldValue, newValue string), size)


	for i := range funcs {
		logger.Infof("Creating function at position %v", i)
		// Must be reallocated in order to keep i from incrementing in the func below
		newIndex := i
		funcs[newIndex] = func(oldValue, newValue string) {
			logger.Debugf("Running on position %v - %v", newIndex, &newIndex)
			bools[newIndex] = true
		}

		ss.Sub(&funcs[i])
	}

	ss.Set("a")

	for i, v := range bools {
		if !v{
			t.Errorf("Position %v was not set to true!", i)
		}
	}
}

func TestStringName(t *testing.T) {
	name := "testname123"
	is := NewString(name)
	if is.Name() != name {
		t.Errorf("Name not as expected!")
	}
}

func TestStringUnsub(t *testing.T){
	is := NewString("int1")
	counter := 0

	f1 := func(oldValue, newValue string) {
		i, _ := strconv.Atoi(newValue)
		counter += i
	}

	is.Sub(&f1)
	is.Set("1")
	if counter != 1{
		t.Error("Inner callback function not run")
	}

	is.Unsub(&f1)
	is.Set("2")
	if counter != 1{
		t.Error("Counter should still be 1, callback should have been unregistered!")
	}

}

func TestStringOverrideFunc(t *testing.T) {
	is := NewString("int1")
	counter := 0

	f1 := func(oldValue, newValue string) {
		i, _ := strconv.Atoi(newValue)
		counter += i
	}

	is.Sub(&f1)
	is.Set("1")

	if counter != 1 {
		t.Error("f1 should have been overridden with its 2nd form that subtracts!")
	}


	f1 = func(oldValue, newValue string) {
		i, _ := strconv.Atoi(newValue)
		counter -= i
	}
	is.Sub(&f1) //No effect actually
	is.Set("1")

	if counter != 0 {
		t.Error("f1 should have been overridden with its 2nd form that subtracts!")
	}
}