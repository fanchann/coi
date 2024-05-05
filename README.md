# COI
![coi](/assets/coi.gif)

## How COI works?
![coi_svg](/assets/coi_works.svg)

## Example
add `// coi` tag, before interface

`foobar.go`
```go
package foobar

// coi <- tag
type IFoobar interface{
    GetFoo() (string,error) // better use 2 return or more, if you set 1 return, coi can't parse :'( 
    GetBar() (string,error)
}
```

then generate the code with command: `coi -f foobar.go`, will generate go file with the code\
generated file -->`foo_bar_impl.go`:

```go
package foobar

type newFooBarImpl struct{}

func NewFooBarImpl() FooBar {
	return &implzBu{}
}

type implzBu newFooBarImpl

func (i *implzBu) GetFoo() (string,error) {
	// dont forget to implement!
	panic("not implemented!")
}

func (i *implzBu)  GetBar() (string,error) {
	// dont forget to implement!
	panic("not implemented!")
}
```

## install coi
```sh
go install https://github.com/fanchann/coi@latest
```

## author
[fanchann](https://github.com/fanchann)