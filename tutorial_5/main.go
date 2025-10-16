// Struct
// struct under struct
// struct function
// interface
// anonymous struct
// interface passing function argument

package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type gasEngine2 struct {
	mpg     uint8
	gallons uint8
	owner
	int
}

type electricEngine struct {
	mpkh uint8
	kwh  uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg

}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{name: "Alex"}}
	myEngine = gasEngine{25, 15, owner{"Alex"}}
	myEngine.mpg = 22
	fmt.Printf("struct in struct mpg %v gallons %v struct_owner %v  \n", myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)

	var myEngine2 gasEngine2 = gasEngine2{25, 15, owner{"Alex"}, 33}
	fmt.Printf("struct in struct2 mpg %v gallons %v struct_owner %v int %v \n", myEngine2.mpg, myEngine2.gallons, myEngine2.name, myEngine2.int)

	fmt.Printf("Struct Func, Miles Left %v\n", myEngine.milesLeft())

	anonymousStructEgasEnginexample()

	canMakeIt(myEngine, 14)
	canMakeIt(electricEngine{70, 11}, 14)

}

func anonymousStructEgasEnginexample() {
	anonymousStruct := struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Printf("Anonymous struct mpg %v gallons %v \n", anonymousStruct.mpg, anonymousStruct.gallons)
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Makes It")
	} else {
		fmt.Println("Can't make it")
	}
}
