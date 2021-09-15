package maps

import "fmt"

func MapMapMap() {
	fmt.Println("\t---- Maps like dictionary ----\n")

	// map[<key type>]<value type>
	fmt.Println("** map[<type of key>]<type of value> **")
	status := map[int]string{
		200: "OK",
		300: "Permanent Redirect",
		400: "Bad Request",
		500: "Internal Server Error",
	}
	fmt.Printf("% #v\n\n", status)

	// length of map
	fmt.Println("** length of map with len() **")
	l := len(status)
	fmt.Printf("length: %d\n\n", l)

	// Reassign value
	fmt.Println("** Reassign value **")
	fmt.Println("- map_variable[<key>] = <new value>")
	status[200] = "Okay"
	fmt.Printf("% #v\n\n", status)

	// Insert value to map
	fmt.Println("** insert value **")
	fmt.Println("- map_variable[<new key>] = <value>")
	status[285] = "เมาแน่เลย"
	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n\n", len(status))

	// Read value in map
	fmt.Println("** Read value **")
	fmt.Println("- variable = map_variable[<key>]")
	value := status[200]
	fmt.Printf("%#v\n\n", value)

	// Delete value in map
	fmt.Println("** delete value **")
	fmt.Println("- use delete(<map_variable>, <key>)")
	delete(status, 285)
	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n\n", len(status))

	// Read value but no key in map
	fmt.Println("** Read value but no key **")
	fmt.Println("- Return empty value or zero value of type")
	v := status[666]
	fmt.Printf("%#v\n\n", v)

	// Check key in map
	fmt.Println("** Check key **")
	if v, flag := status[666]; flag {
		fmt.Printf("value : %q \n", v)

	} else {
		fmt.Println("key not found")
	}
	fmt.Println()

	// Zero value of map
	fmt.Println("** Zero value; Beware!! **")
	var m map[string]string

	if m == nil {
		fmt.Printf("m is nil. value : %#v\n", m)

	} else {
		fmt.Printf("m is not nil. value : %#v\n", m)
	}

	// Allocate memory
	fmt.Println("** allocate memory **")
	fmt.Println("- use make(map[<key type>]<value type>)")
	alloc_m := make(map[string]string)
	fmt.Printf("alloc_m : %#v\n\n", alloc_m)
}
