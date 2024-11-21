package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file "

	file, err := os.Create("./mylcogofile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./mylcogofile.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("databyte is:\n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
