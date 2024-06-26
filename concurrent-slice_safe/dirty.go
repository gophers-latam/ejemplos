package main

import (
	"fmt"
	"time"
)

var shared []int = []int{1, 2, 3, 4, 5, 6}

func main() {
	for i := 0; i < 5; i++ {
		go increase(i)
	}
	for i := 0; i < 5; i++ {
		go decrease(i)
	}
	time.Sleep(2 * time.Second)
}

// increase each element by 1
func increase(num int) {
	fmt.Printf("[+%d a] : %v\n", num, shared)
	for i := 0; i < len(shared); i++ {
		time.Sleep(20 * time.Microsecond)
		shared[i] = shared[i] + 1
	}
	fmt.Printf("[+%d b] : %v\n", num, shared)
}

// decrease each element by 1
func decrease(num int) {
	fmt.Printf("[-%d a] : %v\n", num, shared)
	for i := 0; i < len(shared); i++ {
		time.Sleep(10 * time.Microsecond)
		shared[i] = shared[i] - 1
	}
	fmt.Printf("[-%d b] : %v\n", num, shared)
}
