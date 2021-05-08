package main

import (
	"fmt"
	"strconv"
	"testing"
)

var testData = []int{10, 11, 17, 10}

// go test -run TestSampleSub -v
func TestSampleSub(t *testing.T) {
	expected := "10"
	for _, val := range testData {
		tc := val
		t.Run(fmt.Sprintf("input = %d", tc), func(t *testing.T) { // sub test
			if expected != strconv.Itoa(tc) {
				t.Fail()
			}
		})
	}
}
