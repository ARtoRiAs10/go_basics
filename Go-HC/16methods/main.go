package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	gaurav := User{"gaurav", "gaurav@go.dev", true, 16}
	fmt.Println(gaurav)
	fmt.Printf("gaurav details are: %+v\n", gaurav)

	fmt.Printf("Name is %v and email  is %v.\n", gaurav.Name, gaurav.Email)

	gaurav.GetStatus()
	gaurav.NewMail()
	fmt.Printf("Name is %v and email  is %v.\n", gaurav.Name, gaurav.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
