package pointer

import "fmt"

func GoPointer() {
	fmt.Println("\t---- Pointer in go ----\n")

	me := "Minnie"
	fmt.Printf("You are %s\n\n", me)

	// Get address
	fmt.Println("- use `&<variable>` for get address")
	addr := &me
	fmt.Println("addr me : ", addr)
	fmt.Printf("type of addr : %T\n", addr)
	fmt.Println("- It's like... var addr *string := &me\n")

	// Reasign value
	fmt.Println("- use `*<variable address>` for access address")
	*addr = "Lester"
	fmt.Printf("You are %s\n\n", me)
}
