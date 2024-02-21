package main

import (
	"executemycode/internal/executer"

	"github.com/google/uuid"
)

func main() {
	// Testing Program Executer
	id := uuid.New()
	language := executer.Golang
	code := `
	package main

import (
	"fmt"
)

func main() {
	var num1 string

	fmt.Println("Enter the first string: ")
	fmt.Scan(&num1)

	// Printing the result
	fmt.Printf("The String You entered is %s", num1)
}
`

	program := executer.New(id, language, code)
	go program.Execute()

}
