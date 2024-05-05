package utils

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/fanchann/coi/internal/types"
)

func ErrorWrapper(err error, s ...string) {

	switch {
	case errors.Is(err, types.ErrInterfaceNotFound),
		errors.Is(err, types.ErrPackageNameNotFound),
		errors.Is(err, types.ErrTagCoiNotFound),
		err != nil:
		log.Printf("[COI] Failed generate: %v \n", err)
		os.Exit(1)
	default:
		msg := strings.Join(s, "")
		log.Printf("[COI] Success generate,saved at %v \n", msg)
	}

}
