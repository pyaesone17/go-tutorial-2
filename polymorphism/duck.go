// "If it walks like a duck and it quacks like a duck, then it must be a duck"

package polymorphism

import "fmt"

type Duck struct {
}

func (d Duck) Quack() {
	fmt.Println("Quck quck I am duck")
}

type Chicken struct {
}

func (d Chicken) Zuack() {
	fmt.Println("Zuack quck I am duck")
}
