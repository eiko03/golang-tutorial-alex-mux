// array
// slice
// map
// for loop
// while loop

package main

import (
	"fmt"
	"time"
)

func main() {
	arrayExample()
	sliceExample()
	mapExample()
	loopExample()
	sliceAdvancedExample()

}

func arrayExample() {
	var intarr [3]int32
	intarr[1] = 123
	fmt.Println(intarr[0])
	fmt.Println(intarr[1:3])

	fmt.Println(&intarr[0])
	fmt.Println(&intarr[1])
	fmt.Println(&intarr[2])

	intarr = [3]int32{1, 2, 3}
	fmt.Println(intarr)

	intarr2 := [...]int32{1, 2, 3}
	fmt.Println(intarr2)

	for i, v := range intarr {
		fmt.Printf("the index is %v and value is %v\n", i, v)
	}
}

func sliceExample() {
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length %v with capacity %v\n", len(intSlice), cap(intSlice))

	intSlice2 := []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Printf("The length %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
	fmt.Printf("The length %v with capacity %v\n", len(intSlice3), cap(intSlice3))

}

func sliceAdvancedExample() {
	n := 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Time without preallocation %v\n", timeloop(testSlice, n))
	fmt.Printf("Time with preallocation %v\n", timeloop(testSlice2, n))

}

func timeloop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)

}

func mapExample() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"Adam": 23, "Eve": 45}
	fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["Json"]

	if ok {
		fmt.Printf("age is %v", age)
	} else {
		fmt.Println("There is no such person")
	}

	delete(myMap, "Adam")

	for name, age := range myMap2 {
		fmt.Printf("%v is %v\n", name, age)
	}
	// map loop order is inconsistent
}

func loopExample() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		if i >= 10 {
			break
		}
		i++
	}

	for i2 := 0; i2 < 10; i2++ {
		fmt.Println(i2)
	}
}
