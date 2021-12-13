package lang

import (
	"fmt"
)

type FunctionExample struct{}

func init() {
	fmt.Println("Function Init")
}

func (FunctionExample) Greetings() string {
	fmt.Println("Function Examples")
	return "Function Example"
}

func (e FunctionExample) Setup() {
	fmt.Println("Setup")
}

func (e FunctionExample) Demonstrate() {
	e.WithNoParameters()

	fmt.Printf("Sum of two numbers : %d\n", e.WithMultipleParameter(3, 5))

	message, result := e.WithMultipleReturnValues(3, 5)
	fmt.Printf("%s %d\n", message, result)

}

func (e FunctionExample) Finish() {
	fmt.Println("Finish")
}

func (e FunctionExample) WithNoParameters() {
	fmt.Println("Function call with no parameters")
}

func (e FunctionExample) WithMultipleParameter(a int, b int) int {
	fmt.Println("Function call with multiple parameters")
	return a + b
}

func (e FunctionExample) WithMultipleReturnValues(a int, b int) (string, int) {
	fmt.Println("Function call with multiple return parameters")
	return "Sum of two number is ", a + b
}
