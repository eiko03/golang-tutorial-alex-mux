// strings
// utf8
// stringbuilder
// rune

package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	stringUtf8Example(myString)

	myString2 := []rune("résumé")
	runeUtf8Example(myString2)

	runeExample()

	stringBuildingExample()

}

func stringBuildingExample() {
	var strSlice = []string{"a", "b", "c", "d", "e"}
	var catstr = ""
	for i := range strSlice {
		catstr += strSlice[i]
	}
	fmt.Printf("built via concat %v\n", catstr)

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	catstr = strBuilder.String()
	fmt.Printf("built via strBuilder %v\n", catstr)

	// strings are immutable
	// string builder is faster
}

func stringUtf8Example(myString string) {

	var indexed = myString[0]
	fmt.Printf("The first byte: %v %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Printf("index is %v and value is %v\n", i, v)
	}

	fmt.Printf("len of this string is %v \n", len(myString))
}

func runeUtf8Example(myString []rune) {

	var indexed = myString[0]
	fmt.Printf("The first byte: %v %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Printf("index is %v and value is %v\n", i, v)
	}

	fmt.Printf("len of this string is %v \n", len(myString))
}

func runeExample() {
	var myRune = 'a'
	fmt.Printf("my Rune = %v \n", myRune)
}
