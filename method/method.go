package method

import "fmt"

type User struct {
	Username string
	FullName string
	Email    string
}

// info() is not function, it is `method`
// 'cause >> (u User) << is `behavior` of struct
// (u User) is Receiver function
func (u User) Info() {
	fmt.Printf("User name	: %v\n", u.Username)
	fmt.Printf("Full name	: %v\n", u.FullName)
	fmt.Printf("Email		: %v\n", u.Email)
}

func BasicMethod() {
	fmt.Println("\t---- Method ----\n")
	u := User{
		Username: "dolly",
		FullName: "Dolly Radionoise",
		Email:    "dolly0@sister.com",
	}

	// Call method
	fmt.Println("** call method **")
	fmt.Println("- use instance struct for call method of struct")
	fmt.Println("\tLike... u.info()")
	u.Info()

	fmt.Println("\n\nTips:")
	fmt.Println("`Uppercase` name such as variable, function, struct or method")
	fmt.Println("\tIt's `Public`")
	fmt.Println("And `Lowwercase` is `Private`")
}
