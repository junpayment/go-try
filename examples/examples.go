package main

import (
	try "gotry"
	"log"
)

func main() {
	f := func() (v int, e error) {
		return 1, nil
	}

	defer func() {
		err := recover()
		if err == nil {
			return
		}
	}()
	res := try.Try(f())
	log.Println(res)
}
