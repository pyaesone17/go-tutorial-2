package pointer

import "fmt"

// Point ...
func Point() {

	var foo = 1
	fmt.Println(foo, "Foo")

	var bar = &foo

	fmt.Println(bar, "Bar")

	*bar = 14
	fmt.Println(*bar)
}
