package main

import (
	"nyan/encapsulation/auth"
)

func main() {
	// human := &composition.Human{}
	// superman := &composition.Superman{Human: human}
	// superman.Work()
	// pointer.Point()

	auth.Check("token")

}

func Check() {

}
