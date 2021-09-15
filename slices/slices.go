package slices

import "fmt"

func BasicSlice() {
	fmt.Println("\t---- Slices ----\n")

	langs := []string{"go", "python", "ruby"}
	fmt.Printf("langs: %#v\n", langs)

	n := langs[0]
	fmt.Printf("langs[0]: %#v\n", n)

	langs[1] = "Pythonistas"
	fmt.Printf("langs: %#v\n", langs)

	// Size of Array
	fmt.Println("\n** Size of Array with len() **")
	l := len(langs)
	fmt.Println("length of langs: ", l)

	// Insert value to array
	fmt.Println("\n** Insert data to array **")
	langs = append(langs, "rust", "js", "prolog")
	fmt.Printf("langs: %#v\n", langs)

	// haft open range
	fmt.Println("\n** haft open range **")
	a := langs[0:2]
	fmt.Printf("a	 : %#v\n", a)

	// get first to last index
	b := langs[0:len(langs)]
	c := langs[0:3]
	d := langs[:3]
	e := langs[:]
	fmt.Printf("b	 : %#v\n", b)
	fmt.Printf("c	 : %#v\n", c)
	fmt.Printf("d	 : %#v\n", d)
	fmt.Printf("e	 : %#v\n", e)
	fmt.Printf("langs: %#v\n", langs)

	printSlice(langs)

	// cords is array
	cords := [4]string{"A", "B", "C", "Z"}
	// cords[:] is slice (all elem)
	printSlice(cords[:]) // Can't send array 'cause @param : slice ([]string)
}

func printSlice(ns []string) {
	fmt.Printf("printSlice -> ns: %#v\n", ns)
}
