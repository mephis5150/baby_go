package condition

import (
	"fmt"
	"time"
)

func Conditions() {
	fmt.Println("\t---- Condition ----\n")

	fmt.Println("** If condition **")

	num := 13

	if num%2 == 0 {
		fmt.Println("เลขคู่")

	} else if num == 31 {
		fmt.Println("คนมีคู่ไม่รู้หรอก")

	} else {
		fmt.Println("เลขโสด")
	}

	fmt.Println("\n** Switch...Case **")

	today := time.Friday

	switch today {
	case time.Monday:
		fmt.Println("today is Monday")

	case time.Friday:
		fmt.Printf("Last %v night\n", today)
		fallthrough

	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")

	default:
		fmt.Printf("สวัสดีวัน %v อยู่ดีมีแฮงเดย์\n", today)
	}
}
