package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	gaurav := User{"gaurav", "gaurav@go.dev", true, 16}
	fmt.Println(gaurav)
	fmt.Printf("gaurav details are: %+v\n", gaurav)

	fmt.Println("Name is %v and email is %v.", gaurav.Name, gaurav.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
