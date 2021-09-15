package loop

import "fmt"

func ForBasic() {
	fmt.Println("\t---- Loop in go ----\n")

	// For loop
	fmt.Println("** for loop **")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// While loop
	fmt.Println("** while loop **")
	fmt.Println("- no while loop in go, we use `for`")

	k := 0
	for k < 10 {
		fmt.Println(k)
		k++
	}
	fmt.Println("- forever loop in go, use `for` but no condition\n")

	// loop and slice
	fmt.Println("** loop and slice **")
	langs := []string{"go", "python", "ruby", "rust"}

	for index, value := range langs {
		fmt.Println(index, " : ", value)
	}
}
