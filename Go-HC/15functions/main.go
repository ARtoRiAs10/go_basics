package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeting()

	greetingTwo()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proRes, myMessage := proAdder(2, 7, 8, 9, 22, 6)
	fmt.Println("Pro result is:", proRes)
	fmt.Println("Pro result is:", myMessage)
}

func adder(valOne int, valTwo int) int { //signature whta type value you will return
	return valOne + valTwo
}

func greetingTwo() {
	fmt.Println("Another method")
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi from Pro result function"
}

func greeting() {
	fmt.Println("Namstey from golang")

}
