package b

import (
	"github.com/fnaranjo-vmw/multiversion-companion/a"
)

func NewFoo(data int) a.Foo {
	return a.Foo{Data: data}
}

func DefaultFoo() a.Foo {
	return a.Foo{Data: 992}
}
