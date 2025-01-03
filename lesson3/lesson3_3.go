package	main

import "fmt"

type User struct {
	ID int
	Name string
}

func main() {
	user := User{
		ID: 1,
		Name: "Shishi",
	}

	fmt.Println("ID:", user.ID)
	fmt.Println("Name:", user.Name)
}