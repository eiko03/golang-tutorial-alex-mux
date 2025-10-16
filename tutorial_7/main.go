// go routine
// waitgroup to wait for all routines to finish
// mutex to protect shared resources from concurrent editing (results slice)

package main

import (
	"fmt"
	"sync"
	"time"
)

// var m sync.Mutex = sync.Mutex{}
var m = sync.RWMutex{}
var wg, wg1, wg2 = sync.WaitGroup{}, sync.WaitGroup{}, sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = make([]string, len(dbData))

func main() {
	t0 := time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	for i := 0; i < 2000; i++ {
		wg1.Add(1)
		go dbCall2()
	}
	for i := 0; i < 2000; i++ {
		wg2.Add(1)
		go count()
	}
	wg.Wait()
	wg1.Wait()
	wg2.Wait()
	fmt.Println("total time: ", time.Since(t0))
	fmt.Println("Results: ", results)

}

func count() {
	res := 0
	for i := 0; i < 100000; i++ {
		res += 1
	}
	wg2.Done()
}

func dbCall(i int) {

	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()

}

func dbCall2() {

	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	wg1.Done()

}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("Current Results: %v \n", results)
	m.RUnlock()
}
