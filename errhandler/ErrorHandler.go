package errhandler

import (
	"fmt"
	"log"
)

type ErrorHandler struct {
}

func (this *ErrorHandler) HandleError(err error, isFatal bool) {
	if err != nil {
		if isFatal {
			log.Fatal(err)
		} else {
			fmt.Println(err)
		}
	}
}
