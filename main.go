package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func p1(x int) {
	//Instantiating data types and variables
	n := x
	s := make([]int, n, n+1)
	m := make(map[string]int)
	var a []int
	count := 0
	now := time.Now()

	// Initializing slice and returning time:
	start := now
	for i := range s {
		s[i] = count
		count++
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	// Initializing map and returning time:
	start = now
	for i := range s {
		index := strconv.Itoa(i)
		m[index] = i
	}
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	// Initializing array and returning time:
	start = now
	for _, v := range s {
		a = append(a, v)
	}
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	// Increment each data type return time:
	// Slice
	start = now
	for i := range s {
		s[i]++
	}
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	// Array
	start = now
	for i := range s {
		a[i]++
	}
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	// Map
	start = now
	for i := range m {
		m[i]++
	}
	elapsed = time.Since(start)
	fmt.Println(elapsed)

}

func p2(x int) {
	var arr []int
	slice := make([]int, x, x+1)
	now := time.Now()

	// Generate random numbers and store in data types
	for i := 0; i < x; i++ {
		num := rand.Intn(100)
		slice[i] = num
		arr = append(arr, num)
	}
	// Sorting array
	start := now
	a1 := arr
	sort.Ints(a1)
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	// Sorting slice
	start = now
	s1 := slice
	sort.Ints(s1)
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	fmt.Println(slice, arr)
}

func main() {
	// Get x from command line and handle nil inputs -- taken from stackoverflow
	var x int
	fmt.Println("Enter an integer value: ")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Program 1:")
	p1(x)
	fmt.Println("Program 2:")
	p2(x)
}
