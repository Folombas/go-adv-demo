package main

import "fmt"

type User struct {
	Name string
}
// go run -gcflags '-m -l' main.go
func main() {
	user := &User {
		Name: "Vasya",
	}
	fmt.Println(user)
}
