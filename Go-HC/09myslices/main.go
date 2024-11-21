package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to lecture on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("Type of friutlist is %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 432

	highScores[2] = 562
	highScores[3] = 892
	// highScores[4] = 239

	highScores = append(highScores, 555, 666, 879)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index
	var courses = []string{"reacjs", "javascript", "swift", "python", "ruby"}

	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
