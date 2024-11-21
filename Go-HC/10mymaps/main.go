package main

import "fmt"

func main() {
	fmt.Println("Maps in golang ")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Lis of all Languages:", languages)
	fmt.Println("js is short for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Lis of all Languages:", languages)

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)

	}
}
