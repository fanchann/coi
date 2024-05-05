package main

import "net/http"

type FooBar interface {
	Foo() error
	Bar() error
}

type IWebController interface {
	CreateNewUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}
