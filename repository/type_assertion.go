package repository

import "fmt"

type test struct {
	ID int
}

func run() {
	x := test{1}

	typeAssertion(x)
}

func typeAssertion(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("The string is '%s'.\n", i.(string))
	case int:
		fmt.Printf("argument is int")
	case test:
		fmt.Printf("argument is int")
	default:
		fmt.Printf("Unsupported value (type=%T)\n", i)
	}

}
