package main

import "fmt"

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	right int
}

func main() {
	var u User = User{
		Name:  "polo",
		Email: "polo@polo.polo",
	}
	fmt.Printf("user info : %v\n", u)

	a := Admin{
		right: 2,
	}
	a.Name = "Polette"
	a.Email = "polette@polo.polo"
	fmt.Printf("Admin user = %v\n", a)

}
