package errorx

import (
	"github.com/go-errors/errors"
	"log"
)

// PrintDetailedError prints error detailed trace
func PrintDetailedError(err error) {
	if err != nil {
		errStack, ok := err.(*errors.Error)
		if ok {
			log.Println(errStack.ErrorStack())
		} else {
			log.Println(err)
		}
	}
}
