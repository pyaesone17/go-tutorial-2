package polymorphism

import "fmt"

type Nyan struct {
}

func (nyan Nyan) Quack() {
	fmt.Println("Quck quck but I am Nyan")
}
