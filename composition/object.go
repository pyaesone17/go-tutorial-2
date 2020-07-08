package composition

import "fmt"

type Human struct {
}

func (human Human) Work() {
	fmt.Println("Work")
}

func (human Human) Eat() {
	fmt.Println("Eat")
}

type Superman struct {
	*Human
}

func (superman Superman) Work() {
	fmt.Println("Nice --------")
	superman.Human.Work()
}

func (superman Superman) Attack() {
	fmt.Println("Phew phew")
}
