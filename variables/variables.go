package variables

import "fmt"

// global variable
var (
	StrVar     string = "Type String"
	ByteVar    byte   = 255 // range 0-255
	RuneVar    rune   = 'A'
	BooleanVar bool   = true
)

func IntegerRange() {
	/* ** NOTE **
			__Arithmetic operators__
	กำหนดให้
	dec = 10
	bin = 1010

	# shift bit
	Operator: "right shift" 10 >> 1	// ค่า 10 (1010) right shift 1 bit
	Solve:
		กำหนดให้ * เป็นตัวเริ่มต้น
		right shift 1 bit = |*1|0|1|0| >> |0|*1|0|1|
	answer = 0101 = 5

	# bitwise
	Operator: "XOR" 10 ^ 0 // 10 XOR กับ 0
	Solve:
		|1|0|1|0| ^ |0|0|0|0|	// เหมือน 0 ต่าง 1
	answer = 1010 = 10
	** END NOTE ** */

	// Range of data type in Go
	fmt.Println("\n** Range of uint and int")
	const MaxUint = ^uint(0) // infinite XOR 0 => unsigned type
	const MinUint = 0        // minimum unsigned int = 0

	const MaxUint8 = ^uint8(0)
	const MaxUint16 = ^uint16(0)
	const MaxUint32 = ^uint32(0)
	const MaxUint64 = ^uint64(0)

	const MaxInt = int(MaxUint >> 1) // right shift 1 bit
	const MinInt = -MaxInt - 1       // ค่าสุดท้ายของจำนวนที่มากที่สุด = n-1 ฉันใด ค่าแรกของจำนวนที่น้อยที่สุด = -(จำนวนมากสุด)+(-1) ฉันนั้น

	const MaxInt8 = int8(MaxUint8 >> 1)
	const MinInt8 = -MaxInt8 - 1

	const MaxInt16 = int16(MaxUint16 >> 1)
	const MinInt16 = -MaxInt16 - 1

	const MaxInt32 = int32(MaxUint32 >> 1)
	const MinInt32 = -MaxInt32 - 1

	const MaxInt64 = int64(MaxUint64 >> 1)
	const MinInt64 = -MaxInt64 - 1

	fmt.Println("usigned int:")
	fmt.Printf("\tuint: %d to %d\n", MinUint, MaxUint)
	fmt.Printf("\tuint8: %d to %d\n", MinUint, MaxUint8)
	fmt.Printf("\tuint16: %d to %d\n", MinUint, MaxUint16)
	fmt.Printf("\tuint32: %d to %d\n", MinUint, MaxUint32)
	fmt.Printf("\tuint64: %d to %d\n", MinUint, MaxUint64)

	fmt.Println("int:")
	fmt.Printf("\tint: %d to %d\n", MinInt, MaxInt)
	fmt.Printf("\tint8: %d to %d\n", MinInt8, MaxInt8)
	fmt.Printf("\tint16: %d to %d\n", MinInt16, MaxInt16)
	fmt.Printf("\tint32: %d to %d\n", MinInt32, MaxInt32)
	fmt.Printf("\tint64: %d to %d\n", MinInt64, MaxInt64)
}

func VariableDoc() {
	// String
	email := "mephis2015@protonmail.com"

	// Integer
	var age = 33

	// Float
	const Pi float32 = 3.14

	// The zero value
	var ZeroString string // "" (empty string)
	var ZeroInt int       // 0
	var ZeroFloat float64 // 0
	var ZeroBoolean bool  // false

	// Function scope ใช้ได้แค่ใน function นั้น ๆ
	fmt.Println("** Function scope: ตัวแปรใช้ได้แค่ใน function นั้น ๆ")
	fmt.Println("การประกาศตัวแปร นอกจาก var VAR1 string = \"Something\" ยังสามารถใช้วิธีประกาศตัวแปรภายใน function ได้ดังนี้")
	fmt.Printf("\temail := \"%v\"\n", email) // ตัวแปรใน Go เมื่อประกาศแล้วต้องเรียกใช้
	fmt.Println("\tvar age = ", age)
	fmt.Println("\tvar pi float32 = ", Pi)

	fmt.Println("\n** Zero value is default of each data type. => ค่าเริ่มต้นของตัวแปรที่ไม่ได้กำหนดค่า")
	fmt.Println("Zero value of String : ", ZeroString)
	fmt.Printf("Zero value of String : %q\n", ZeroString) // %q = quotes คำที่ print
	fmt.Println("Zero value of Int : ", ZeroInt)
	fmt.Println("Zero value of Float64 : ", ZeroFloat)
	fmt.Println("Zero value of Boolean : ", ZeroBoolean)

	// Package scope ทั้ง package สามารถใช้ได้
	fmt.Println("\n** Package scope: ตัวแปรสามารถเรียกใช้ได้ทั้ง package (Global variable)")
	fmt.Println("ทำได้โดยการประกาศไว้ก่อนการสร้าง function ดังนี้")
	fmt.Printf("var (\n\tStrVar     string = \"Type String\"\n\tByteVar    byte   = 255\n\tRuneVar    rune   = 'A'\n\tBooleanVar bool   = true\n)\n")
	fmt.Println("Output:")
	fmt.Println("\tstring : ", StrVar)
	fmt.Println("\tbyte (0-255) : ", ByteVar)
	fmt.Println("\trune (convert char to int32) : ", RuneVar, "(A)")
	fmt.Println("\tbool : ", BooleanVar)

	// Constants
	fmt.Println("\n** Constants: การประกาศตัวแปรแบบไม่สามารถกำหนดค่าใหม่ได้")
	fmt.Printf("\tconst Pi float32 = %v\n", Pi)

	// iota การกำหนดค่าต่อไปแบบอัตโนมัติ เริ่มจาก 0 และ +1 ไปเรื่อย ๆ
	fmt.Println("\n** iota ใช้สำหรับการกำหนดค่าแบบอัตโนมัติโดยเริ่มจาก 0 และ +1 ไปเรื่อย ๆ")
	fmt.Println("หากกำหนด iota ให้ตัวแปรตัวที่ 2 ค่าของ iota จะเริ่มที่ 1 เนื่องจากตัวแปรก่อนหน้าเป็น index 0")

	const (
		Zero = 0
		One  = iota
		Two
		Three
		Four
		Five
		Six
		Seven
	)

	fmt.Println("ตัวอย่างการใช้ iota")
	fmt.Println("const (\n\tZero = 0\n\tOne  = iota\n\tTwo\n\tThree\n\tFour\n\tFive\n\tSix\n\tSeven\n)")
	fmt.Println("Output:")
	fmt.Printf("\tZero = %v\n", Zero)
	fmt.Printf("\tOne = %v\n", One)
	fmt.Printf("\tTwo = %v\n", Two)
	fmt.Printf("\tThree = %v\n", Three)
	fmt.Printf("\tFour = %v\n", Four)
	fmt.Printf("\tFive = %v\n", Five)
	fmt.Printf("\tSix = %v\n", Six)
	fmt.Printf("\tSeven = %v\n", Seven)

	fmt.Println("\n*Tips: \n\tการประกาศตัวแปรโดยใช้ := ไม่สามารถใช้ภายนอก function ได้\n\tตัวแปรใน Go เมื่อประกาศแล้วต้องเรียกใช้เท่านั้น")
}
