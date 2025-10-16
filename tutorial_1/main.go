// Varables
// Arithmatic Operations
// Aritmatic Operations
// Concat

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int
	fmt.Println(intNum)

	var floatNum float64
	fmt.Println(floatNum)

	var floatNum32 float32 = 3.14
	var intNum32 int32 = 42
	var res float32 = float32(intNum32) + floatNum32
	fmt.Println(res)

	var intNum1 int = 3
	var intNum2 int = 2

	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var mayString string = "Hello" + " " + "World!"
	fmt.Println(mayString)

	fmt.Println(utf8.RuneCountInString("y"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 int
	fmt.Println(intNum3)

	myVar := "text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "constant"
	fmt.Println(myConst)

	const pi float32 = 3.14
	fmt.Println(pi)

}
