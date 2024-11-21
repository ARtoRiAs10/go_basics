package main

import "fmt"

// var jwtToken int = 300000
const LoginToken string = "Hello" //Public

func main() {
	var username string = "Gaurav"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.4554578
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var smallFloat2 float64 = 255.4554578
	fmt.Println(smallFloat2)
	fmt.Printf("Variable is of type: %T \n", smallFloat2)

	//default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Another variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "Learnit"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Another variable is of type: %T \n", LoginToken)

}
