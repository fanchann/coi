package main

type newFooBarImpl struct{}

func NewFooBarImpl()FooBar {
	return &impleBw{}
}
	
type impleBw newFooBarImpl
	

func (i *impleBw) Foo() error {
	// dont forget to implement!
	panic("not implemented!")
}

func (i *impleBw) Bar() error {
	// dont forget to implement!
	panic("not implemented!")
}
