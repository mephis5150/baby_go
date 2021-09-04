package function

import "fmt"

func Function() {
	fmt.Println("\t---- Function ----\n")

	fmt.Println("** Call function **")
	info("Leskinen", "Alexis", 25)
	fmt.Println()

	fmt.Println("** Function return **")
	day := today()
	fmt.Printf("today return: %v\n", day)
	fmt.Println()

	fmt.Println("** Function return multiple value **")
	a, b, c := swap("A", "B", 33)
	fmt.Printf("Call swap: %s, %s, %d\n", a, b, c)
	fmt.Println()
}

func info(name, msg string, age int) {
	fmt.Println("This is info")

	fmt.Printf("This is %s\n", name)
	fmt.Printf("My message is %s\n", msg)
	fmt.Printf("I'm %d years old\n", age)
}

func today() string {
	return "มื้อนี่"
}

func swap(x, y string, z int) (string, string, int) {
	return y, x, z
}
