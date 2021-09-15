package interfaces

import "fmt"

// Create interface
type Phone interface {
	Call(number string)
}

// Create struct
type Xiaomi struct {
	Name string
}

// Create method
func (x Xiaomi) Call(number string) {
	fmt.Println(x.Name, "calling", number)
}

func (x Xiaomi) Answer() {
	fmt.Println("Good morning, Toon&Tar's Cafe. How may I help you?")
}

// Create function
// @param : p Phone is interface
func Dial(p Phone) {
	p.Call("+66123456789")
}

func BasicInterface() {
	fmt.Println("\t---- Interface ----\n")

	x := Xiaomi{
		Name: "Redmi 9T",
	}

	// Call Dial() function,
	// and use Call() method
	Dial(x)

	// Can not call Answer() method in Dial() function,
	// 'cause Answer() is not in Phone interface.
	// However you can call method directly.
	x.Answer()
}
