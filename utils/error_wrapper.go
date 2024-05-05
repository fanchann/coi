package utils

import (
	"errors"
	"log"
	"os"

	"github.com/fanchann/coi/internal/types"
)

func ErrorWrapper(err error, s ...string) {

	switch {
	case errors.Is(err, types.ErrInterfaceNotFound),
		errors.Is(err, types.ErrPackageNameNotFound),
		errors.Is(err, types.ErrTagCoiNotFound),
		err != nil:
		log.Printf("[COI] generate failed!: %v \n", err)
		os.Exit(1)
	}

}
