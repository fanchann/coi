package main

type newFooBarImpl struct{}

func NewFooBarImpl()FooBar {
	return &imploqO{}
}
	
type imploqO newFooBarImpl
	

func (i *imploqO) Foo() error {
	// dont forget to implement!
	panic("not implemented!")
}

func (i *imploqO) Bar() error {
	// dont forget to implement!
	panic("not implemented!")
}
