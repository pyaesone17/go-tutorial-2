package composition

import (
	"testing"
)

func TestClient(t *testing.T) {

	johnDoe := &Human{}
	super := &Superman{johnDoe}

	super.Eat()
	super.Work()

	super.Attack()
}
