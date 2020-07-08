package polymorphism

import (
	"testing"
)

func TestClient(t *testing.T) {

	nyan := &Nyan{}
	duck := &Duck{}

	process(duck)
	process(nyan)

}
