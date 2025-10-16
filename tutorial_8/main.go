// Go by Example: Channels
// buffered and unbuffered, selection

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_CHICKEN_PRICE float32 = 5
const MAX_TOFU_PRICE float32 = 7

func main() {
	runUnbufferedExample()
	runBufferedExample()
	runChickenTofuExample()

}

func runChickenTofuExample() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"wallmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrice(websites[i], chickenChannel)
		go checkTofuPrice(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrice(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}
func checkTofuPrice(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-tofuChannel:
		fmt.Println("Found a deal on tofu at", website)
	case website := <-chickenChannel:
		fmt.Println("Found a deal on chicken at", website)
	}

}

func runUnbufferedExample() {
	var unbuffered_channel = make(chan int)
	go unbuffered_channel_process(unbuffered_channel)
	for i := range unbuffered_channel {
		fmt.Println("Received unbuffered:", i)
	}
}

func runBufferedExample() {
	var buffered_channel = make(chan int, 5)
	go buffered_channel_process(buffered_channel)
	for i := range buffered_channel {
		fmt.Println("Received buffered:", i)
		time.Sleep(time.Second * 1)
	}
}

func unbuffered_channel_process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}

	fmt.Println("All data sent to unffered channel")

}

func buffered_channel_process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("All data sent to buffered channel")
}
