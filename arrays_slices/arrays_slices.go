package main

import (
	"fmt"
	"sort"
)

func passAnArray(a [3]int) {
	a[0] += 100
}

func passASlice(a []int) {
	if len(a) == 0 {
		a = append(a, 0)
	}
	a[0] += 100
}

func passAnArrayAndReturn(a [3]int) [3]int {
	a[0] += 100
	return a
}

func funFact1() {
	fmt.Println("========================================")
	fmt.Println("Arrays and Slices as function parameters")

	aa := [3]int{1, 2, 3} // 3 element array
	ab := []int{1, 2, 3}  // 3 element slice

	passAnArray(aa)
	passASlice(ab)

	fmt.Println("aa: ", aa) // [1 2 3]
	fmt.Println("ab: ", ab) // [101 2 3]

	aa = passAnArrayAndReturn(aa)
	fmt.Println("aa: ", aa) // [101 2 3]
}

func funFact2() {
	fmt.Println("===========================================================")
	fmt.Println("Slices are doubled in size once they outgrow their capacity")

	d := make([]int, 0)
	fmt.Printf("%T %v %d %d \n", d, d, len(d), cap(d))
	// this will throw an error, there is no position 0 yet
	// d[0] = 10

	d = make([]int, 5) // same as d = make([]int, 5, 5)
	d[0] = 10          // [10 0 0 0 0]
	fmt.Printf("%T %v %d %d \n", d, d, len(d), cap(d))

	d = append(d, 1)
	fmt.Printf("%T %v %d %d \n", d, d, len(d), cap(d))

	d = append(d, 2, 3, 4, 5)
	fmt.Printf("%T %v %d %d \n", d, d, len(d), cap(d))

	d = append(d, 6)
	fmt.Printf("%T %v %d %d \n", d, d, len(d), cap(d))
}

func funcFact3() {
	fmt.Println("=========================================================")
	fmt.Println("The sort package can be used to sort slices & collections")

	// sort in ascending order
	s := []int{5, 3, 7, 2, 4, 1, 6, 9, 8, 10}
	sort.Ints(s)
	fmt.Printf("%T %v\n", s, s)

	// sort in descending order
	s = []int{5, 3, 7, 2, 4, 1, 6, 9, 8, 10}
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	fmt.Printf("%T %v\n", s, s)

	// using sort to sort an array
	a := [10]int{5, 3, 7, 2, 4, 1, 6, 9, 8, 10}
	sort.Ints(a[:])
	fmt.Printf("%T %v\n", a, a)
}

func main() {
	// Checking how arrays and slices are passed in functions
	funFact1()

	// Slice arrays are doubled in size once the slice outgrows
	// the capacity of the underlying array
	funFact2()

	// The sort package can be used to sort slices of integers
	funcFact3()
}
