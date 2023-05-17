package main

import (
	psm "a/b/c/kitex_gen/psm/validator"
	"log"
)

func main() {
	svr := psm.NewServer(new(ValidatorImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
