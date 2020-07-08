package method

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	johnDoe := &Human{name: "John Doe"}
	johnDoe.Hi()

	johnDoe.SetAge(25)

	fmt.Println(johnDoe.age, "AGE")
}
