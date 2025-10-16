// memory location
// pointer
// slice pointerExample
// pointer passing as function argument example

package main

import "fmt"

func main() {

	pointerBasicExample()
	slicePointerExceptionExample()
	pointerPassingToFunctionExample()

}

func pointerPassingToFunctionExample() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("memory location of thing1 array %p\n", &thing1)

	square(&thing1)
	fmt.Printf("THE thing1 explicit pointer func: %v\n", thing1)

	thing1 = [5]float64{1, 2, 3, 4, 5}

	square_simple_for_intelligent_range(&thing1)
	fmt.Printf("THE thing1 explicit pointer func: %v\n", thing1)
	//square and  square_simple_for_intelligent_range is same
}

func square(thing2 *[5]float64) [5]float64 {

	fmt.Printf("memory location of thing2 array %p\n", thing2)
	for i := range *thing2 {
		fmt.Printf("loop %v\n", (*thing2)[i])
		(*thing2)[i] = (*thing2)[i] * (*thing2)[i]

	}
	return *thing2
}

func square_simple_for_intelligent_range(thing2 *[5]float64) [5]float64 {

	fmt.Printf("memory location of thing2 array simple range %p\n", thing2)
	for i := range thing2 {
		fmt.Printf("loop %v\n", thing2[i])
		thing2[i] = thing2[i] * thing2[i]

	}
	return *thing2
}

func pointerBasicExample() {
	var p *int32 = new(int32) //initialize pointer
	var i int32

	fmt.Printf("value of p point to is (derefference) %v \n", *p)
	fmt.Printf("value of i  is  %v \n", i)

	p = &i
	*p = 1 // changes value of i
	fmt.Printf("value of  i: %v \n", i)

	var k int32 = 2
	i = k // variables are copied by value
	fmt.Printf("pointer(i): %p and i: %p k: %p \n", p, &i, &k)
}

func slicePointerExceptionExample() {
	slice := []int8{1, 2, 3, 4}
	sliceCopy := slice
	sliceCopy[2] = 55
	fmt.Printf("slice: %v  copied: %v \n", slice, sliceCopy) //slices are copied by refference
}
