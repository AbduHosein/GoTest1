package main

import (
	"fmt"
	"strconv"
	"time"
)

func p1(x int) {
	//Instantiating data types
	n := x
	s := make([]int, n, n+1)
	m := make(map[string]int)
	var a [20]int
	count := 0
	now := time.Now()

	// Initializing slice and returning t it takes
	before1 := now
	for i, _ := range s {
		s[i] = count
		count++
	}
	fmt.Println(s)
	after1 := now
	time1 := after1.Sub(before1)
	fmt.Println(before1, after1)
	fmt.Println(time1)

	// Initializing map and returning t
	before2 := now
	for i, _ := range s {
		index := strconv.Itoa(i)
		m[index] = i
	}
	after2 := now
	time2 := after2.Sub(before2)
	fmt.Println(time2)

	// Initializing array and returning t
	before3 := now
	for i, v := range s {
		a[i] = v
	}
	after3 := now
	time3 := after3.Sub(before3)
	fmt.Println(time3)

	// Increment each data type
	before := now
	for i, _ := range s {
		s[i]++
	}
	after := now
	t := after.Sub(before)
	fmt.Println(t)

	before = now
	for i, _ := range s {
		a[i]++
	}
	after = now
	t = after.Sub(before)
	fmt.Println(t)

	before = now
	for i, _ := range m {
		m[i]++
	}
	after = now
	t = after.Sub(before)
	fmt.Println(t)
}

func p2(x int) {

}

func main() {
	// Get x from command line and handle nil inputs -- taken from stackoverflow
	var x int
	fmt.Println("Enter an integer value : ")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println(err)
	}
	p1(x)

}
