package b

import (
	"testing"
)

func TestDefaultFoo(t *testing.T) {
	foo := DefaultFoo()
	if foo.Data != 992 {
		t.Fail()
	}
}

func TestNewFoo(t *testing.T) {
	foo := NewFoo(123)
	if foo.Data != 123 {
		t.Fail()
	}
}
