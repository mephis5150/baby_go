package structure

import "fmt"

type User struct {
	Username      string
	FullName      string
	ProfilePicUrl string
}

func Structure() {
	fmt.Println("\t---- Struct ----\n")

	username := "lester"
	fullName := "Leskinen Alexis"
	profilePicUrl := "https://knowhere.fake/gopher.jpg"

	// Use struct
	u := User{}
	fmt.Printf("%#v\n", u)
	fmt.Println()

	// Assign value to struct
	fmt.Println("** assign value **")
	fmt.Println("- user instance.field like u.Username")
	u.Username = username
	u.FullName = fullName
	u.ProfilePicUrl = profilePicUrl
	fmt.Printf("%#v\n", u)
	fmt.Println()

	// Read field in struct
	fmt.Println("** read field **")
	name := u.Username
	fmt.Printf("name = %v\n\n", name)

	// Other way to assign value to struct
	fmt.Println("** The Other way for assign value **")
	s := User{
		Username:      "mephie",
		FullName:      "Mephisto Pheles",
		ProfilePicUrl: profilePicUrl,
	}
	fmt.Printf("%#v\n", s)
}
