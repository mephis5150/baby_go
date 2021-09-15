package array

import "fmt"

func showAll(ns [4]string) {
	fmt.Printf("show all : %#v\n", ns)
}

func BasicArray() {
	fmt.Println("\t---- Array ----\n")

	var langs = [4]string{"go", "python", "ruby"}
	fmt.Printf("langs: %#v\n", langs)

	showAll(langs)

	cords := [4]string{"A", "B", "C"}
	showAll(cords)
}
