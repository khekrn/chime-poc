package main

import (
	"fmt"

	"github.com/khekrn/core/base"
	"github.com/khekrn/core/log"
)

func main() {
	fmt.Println("Hello Go package")
	data := map[string]string{}

	base.NewSuccessResponse("ACCEPTED", data)
	log.Info("Hello Go package")
}
