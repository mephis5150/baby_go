package function

import (
	"fmt"
	"strings"
)

var name = "Minnie"

var plus = func(a int, b int) int { return a + b }

func ValueFunction() {
	fmt.Println("\t---- Function is value ----\n")

	fmt.Printf("value: %v\n", name)
	fmt.Printf("type: %T\n", name)
	fmt.Println()

	fmt.Println("** Call function **")
	sum := plus(1, 1)
	fmt.Println("1 + 1 = ", sum)
	fmt.Printf("type of plus: %T\n", plus)
	fmt.Println()

	fmt.Println("** Call function in function **")
	minus := func(x, y int) int { return x - y }
	fmt.Println("3 - 0 = ", minus(3, 0))
	fmt.Printf("type of minus: %T\n", minus)
	fmt.Println()

	fmt.Println("** Function value is parameter **")
	uppercase := func(name string) string { return strings.ToUpper(name) }
	// fmt.Println(uppercase(name))
	say(uppercase)
}

func say(conv_upper func(string) string) {
	fmt.Println("-->> Call function with param <<--")

	conv_name := conv_upper(name)
	fmt.Println("Ahoy!, ", conv_name)
}
