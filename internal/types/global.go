package types

import "errors"

var (
	ErrPackageNameNotFound = msgError("package name not found in file")
	ErrInterfaceNotFound   = msgError("interface not found")
	ErrTagCoiNotFound      = msgError("tag `// coi` not found!")
)

func msgError(s string) error {
	return errors.New(s)
}
