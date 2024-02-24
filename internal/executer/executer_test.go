package executer

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func SetupProgram(code string) *Program {
	id := uuid.New()
	langauge := Golang
	program := NewProgram(id, langauge)
	program.SetCode(code)
	return program
}

func TestExecuter__HelloWorld(t *testing.T) {
	program := SetupProgram(`
	package main
	func main() {
    	print("Hello World")
	}
	`)
	program.Execute()
}

func TestExecuter__MultiOutput(t *testing.T) {
	program := SetupProgram(`
	package main

	import (
		"fmt"
	)

	func main() {
    	fmt.Println("Hello World")
		fmt.Println("Second Hello World")
		fmt.Println("Third Output is This")
	}
	`)
	program.Execute()
}

func TestExecuter__SingleInput(t *testing.T) {
	program := SetupProgram(`
	package main

	import (
		"fmt"
	)

	func main() {
		var input string 
    	fmt.Println("Enter some string ")
		fmt.Scanln(&input)
		fmt.Printf("User Entered this Input: %s",input)
	}
	`)

	go program.Execute()

	time.Sleep(4 * time.Second)
	program.InputChan <- "Umair is the String"

	time.Sleep(4 * time.Second)
}

func TestExecuter__MultiInput(t *testing.T) {
	program := SetupProgram(`
	package main

	import (
		"fmt"
	)

	func main() {
		var input1, input2 int 
    	fmt.Println("Enter first number ")
		fmt.Scanln(&input1)
		fmt.Println("Enter second number ")
		fmt.Scanln(&input2)
		fmt.Printf("User Entered these Inputs: %d, %d",input1, input2)
		sum := input1 + input2
		fmt.Printf("The sum of the inputs is %d\n",sum)
	}
	`)

	go program.Execute()

	time.Sleep(4 * time.Second)
	program.InputChan <- 10
	program.InputChan <- 5

	time.Sleep(4 * time.Second)
}

func SetupPythonProgram(code string) *Program {
	id := uuid.New()
	langauge := Python3
	program := NewProgram(id, langauge)
	program.SetCode(code)
	return program
}
