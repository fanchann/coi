package generator

import (
	"math/rand"
	"strings"
	"text/template"
	"time"

	"github.com/fanchann/coi/internal/types"
)

func GenerateImplementationContract(i *types.Interfaces) string {
	data := struct {
		PackageName string
		IName       string
		UniqID      string
		IMethod     []string
	}{
		PackageName: i.PackageName,
		IName:       i.IName,
		UniqID:      generateRandomString(),
		IMethod:     i.IMethod,
	}

	// Execute template
	var result strings.Builder
	tmpl := template.Must(template.New("").Parse(types.CodeTemplate))
	if err := tmpl.Execute(&result, data); err != nil {
		panic(err)
	}

	return result.String()
}

func generateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 3)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
