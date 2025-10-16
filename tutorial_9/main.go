// generics
// any type
// multiple types
// constraints
// structs with generics
// json with generics

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contractInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	multipleGenericTypesExample()
	anyGenericTypesExample()
	jsonExample()
	structExample()

}

func structExample() {
	var gasCar = car[gasEngine]{
		"Toyota",
		"Camry",
		gasEngine{
			15,
			30,
		},
	}

	var electricCar = car[electricEngine]{
		"Tesla",
		"Model 3",
		electricEngine{
			75,
			4,
		},
	}

	fmt.Printf("gasCar: %+v\n", gasCar)
	fmt.Printf("electricCar: %+v\n", electricCar)

}

func jsonExample() {
	var contracts []contractInfo = loadJSON[contractInfo]("tutorial_9/contractInfo.json")
	fmt.Printf("contracts: %+v\n\n", contracts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("tutorial_9/purchaseInfo.json")
	fmt.Printf("purchases: %+v\n\n", purchases)
}

func loadJSON[T contractInfo | purchaseInfo](filePath string) []T {
	var data, err = os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

func multipleGenericTypesExample() {
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice[int](intSlice))

	floatSlice := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(sumSlice[float32](floatSlice))
}

func anyGenericTypesExample() {
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(isEmpty(intSlice))

	fmt.Println(isEmpty([]float32{}))

}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum

}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0

}
