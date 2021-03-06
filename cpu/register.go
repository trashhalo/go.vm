// This file contains the implementation of the registers the
// CPU contains.

package cpu

import (
	"fmt"
	"os"
)

// Object is the interface for something we store in a register.
//
// This is done to allow future expansion, and also for neatness.
type Object interface {
	Type() string
}

// IntegerObject is an object holding an integer-value.
type IntegerObject struct {
	Value int
}

// Type returns `int` for IntegerObjects.
func (i *IntegerObject) Type() string { return "int" }

// StringObject is an object holding a string-value.
type StringObject struct {
	Value string
}

// Type returns `string` for StringObjects.
func (i *StringObject) Type() string { return "string" }

// Register holds the contents of a single register, as an object.
//
// This means it can hold either an IntegerObject, or a StringObject.
type Register struct {
	o Object
}

// NewRegister is the constructor for a register.
func NewRegister() *Register {
	r := &Register{}
	r.SetInt(0)
	return (r)
}

// GetInt retrieves the integer content of the given register.
// If the register does not contain an integer that is a fatal error.
func (r *Register) GetInt() int {
	switch arg := r.o.(type) {
	case *IntegerObject:
		return arg.Value
	default:
		fmt.Printf("BUG: Attempting to call GetInt on a register holding a non-integer value: %v\n", r.o)
		os.Exit(3)
	}
	// Not reached
	return 0
}

// SetInt stores the given integer in the register.
func (r *Register) SetInt(v int) {
	r.o = &IntegerObject{Value: v}
}

// GetInt retrieves the string content of the given register.
// If the register does not contain a string that is a fatal error.
func (r *Register) GetString() string {
	switch arg := r.o.(type) {
	case *StringObject:
		return arg.Value
	default:
		fmt.Printf("BUG: Attempting to call GetString on a register holding a non-string value: %v\n", r.o)
		os.Exit(3)
	}

	// Not reached
	return ""
}

// SetString stores the supplied string in the register.
func (r *Register) SetString(v string) {
	r.o = &StringObject{Value: v}
}

// Return the type of a registers contents `int` vs. `string`.
func (r *Register) Type() string {
	return (r.o.Type())
}
