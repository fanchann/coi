package types

const (
	CodeTemplate = `package {{.PackageName}}

type new{{.IName}}Impl struct{}

func New{{.IName}}Impl(){{.IName}} {
	return &impl{{.UniqID}}{}
}
	
type impl{{.UniqID}} new{{.IName}}Impl
	
{{range .IMethod}}
func (i *impl{{$.UniqID}}) {{.}}() {
	// dont forget to implement!
	panic("not implemented!")
}
{{end}}`
)
