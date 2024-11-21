package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d:=0; d< len(days); d++{
	// 	fmt.Println(days[d])

	// }

	// for i := range days {
	// 	fmt.Println((days[i]))
	// }

	// for _, day := range days {
	// 	fmt.Println("index is and value  is \n", day)

	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {
			// break
			rougeValue++
			continue
		}

		fmt.Println("Values is:", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at ")
}
