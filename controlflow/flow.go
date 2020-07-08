package controlflow

import "fmt"

func Flow() {

	// IF else flow
	if true {
		fmt.Println("It is true")
	} else {
		fmt.Println("It is false")
	}

	// For loop
	for i := 0; i < 10; i++ {

	}

	j := 10
	// while loop
	for j > 0 {
		j--
	}

	data := []int{1, 2, 3, 4}

	for _, d := range data {
		fmt.Println(d)
	}
}
