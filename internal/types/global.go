package types

import "errors"

var (
	ErrPackageNameNotFound = errors.New("package name not found in file")
	ErrInterfaceNotFound   = errors.New("interface not found")
	ErrTagCoiNotFound      = errors.New("tag `// coi` not found!")
)
