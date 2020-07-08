package method

import "fmt"

type Human struct {
	name string
	age  int
}

func (human Human) Hi() {
	fmt.Println(human.name + " is Hi")
}

func (human *Human) SetAge(age int) {
	human.age = age
}
