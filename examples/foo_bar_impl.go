package main

type newFooBarImpl struct{}

func NewFooBarImpl()FooBar {
	return &implZZR{}
}
	
type implZZR newFooBarImpl
	

func (i *implZZR) Foo() error {
	// dont forget to implement!
	panic("not implemented!")
}

func (i *implZZR) Bar() error {
	// dont forget to implement!
	panic("not implemented!")
}
