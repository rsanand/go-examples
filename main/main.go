package main

import (
	"examples/lang"
	"fmt"
)

func init() {
	fmt.Println("Main Init")
}

func greetings() {
	fmt.Println("Hello There! Welcome to go lang examples")
}

type Example interface {
	Greetings() string
	Setup()
	Demonstrate()
	Finish()
}

func lifecycle(example Example) {
	fmt.Println("===================================")
	exampleId := example.Greetings()

	fmt.Println("===================================")
	fmt.Printf("Calling %s Setup()\n", exampleId)
	example.Setup()
	fmt.Printf("Calling %s Demonstrate()\n", exampleId)
	example.Demonstrate()
	fmt.Printf("Calling %s Finish()\n", exampleId)
	example.Finish()
}

func main() {
	greetings()
	function := lang.FunctionExample{}
	lifecycle(function)
}
