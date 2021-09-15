package slices

import "fmt"

func ZeroSlice() {
	fmt.Println("\t---- Slicing zero ----\n")

	fmt.Println("** zero value of slice is `nil` **")
	var langs []string
	fmt.Printf("langs : %#v\n", langs)

	if langs == nil {
		fmt.Println(`Yes nil "langs" is a nil slices`)

	} else {
		fmt.Printf("langs is not nil, value: %#v\n", langs)
	}
	fmt.Println()

	// empty slice
	fmt.Println("** Empty slice **")
	emp := []string{}
	fmt.Printf("emp: %#v\n", emp)

	if emp == nil {
		fmt.Println(`Yes nil "emp" is a nil slices`)

	} else {
		fmt.Printf("emp is not nil, value: %#v\n", emp)
	}
}
