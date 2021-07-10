package main

import "fmt"

type User struct {
	Name    string
	Surname string
	Age     int
}

//default and empty constructor
func NewUser() *User {
	user := new(User)
	return user
}

func NewUserWithParams(name, surname string, age int) *User {
	user := new(User)
	user.Name = name
	user.Surname = surname
	user.Age = age
	return user
}

func main() {

	tuna := User{Name: "Tuna"}

	fmt.Println(tuna.Name)

	cemre := new(User)
	cemre.Name = "Cemre"
	fmt.Println(cemre.Name)

	//empty constructor
	sadettin := NewUser()
	sadettin.Age = 21
	fmt.Println(sadettin.Age)

	//params constructor
	ali := NewUserWithParams("Ali", "Deniz", 23)
	fmt.Println(*ali)
}
