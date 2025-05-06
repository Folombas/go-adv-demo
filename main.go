package main

type User struct {
	Name string
}
// go run -gcflags '-m -l' main.go
func main() {
	age := getAge()
	canDrink(age)
}

func canDrink(age *int) bool {
	return *age >= 18
}

func getAge() *int {
	age := 18
	return &age
}