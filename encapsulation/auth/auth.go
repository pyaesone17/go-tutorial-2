package auth

import "fmt"

var PublicVar string

type Human struct {
}

type OkI interface{}

// Check ...
func Check(token string) bool {
	decodedString := decode(token)
	return decodedString == "1"
}

func decode(token string) string {
	fmt.Println("Doing decoded")
	return token
}
