package main

type newFooBarImpl struct{}

func NewFooBarImpl()FooBar {
	return &impldQo{}
}
	
type impldQo newFooBarImpl
	

func (i *impldQo) Foo() error {
	// dont forget to implement!
	panic("not implemented!")
}

func (i *impldQo) Bar() error {
	// dont forget to implement!
	panic("not implemented!")
}
