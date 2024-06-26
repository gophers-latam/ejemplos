package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go increase(i)
	}
	for i := 0; i < 5; i++ {
		go decrease(i)
	}
	time.Sleep(2 * time.Second)
}

var (
	share []int = []int{1, 2, 3, 4, 5, 6}
	mutex sync.Mutex
)

// increase each element by 1
func increaseWithMutex(num int) {
	mutex.Lock()
	fmt.Printf("[+%d a] : %v\n", num, share)
	for i := 0; i < len(share); i++ {
		time.Sleep(20 * time.Microsecond)
		share[i] = share[i] + 1
	}
	fmt.Printf("[+%d b] : %v\n", num, share)
	mutex.Unlock()
}

// decrease each element by 1
func decreaseWithMutex(num int) {
	mutex.Lock()
	fmt.Printf("[-%d a] : %v\n", num, share)
	for i := 0; i < len(share); i++ {
		time.Sleep(10 * time.Microsecond)
		share[i] = share[i] - 1
	}
	fmt.Printf("[-%d b] : %v\n", num, share)
	mutex.Unlock()
}
