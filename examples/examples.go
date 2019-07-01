package main

import (
	"log"

	"github.com/junpayment/gotry"
)

func main() {
	f := func() (v, x int, e error) {
		return 1, 2, nil
	}

	defer func() {
		err := recover()
		if err == nil {
			return
		}
	}()
	res := gotry.Try(f())
	log.Println(res)
}
