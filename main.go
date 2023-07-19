package main

import (
	"qnurye/qnurye.me/routers"
)

func main() {
	r := routers.SetUp()

	err := r.Run()

	if err != nil {
		return
	}
}
