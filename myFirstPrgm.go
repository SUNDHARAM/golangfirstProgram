package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("I love Ammu")
	var a int = 10
	fmt.Printf("%d\n", a)

	var arr [3]int
	arr = [...]int{1, 2, 3}
	printArray(arr[:])

	var b string = "100"
	var c int = 20
	sh, err := strconv.Atoi(b)
	var sha string = strconv.Itoa(c)
	if err != nil {
		fmt.Printf("%v\n", sh)
	}
	fmt.Printf("%v\n", sha)

	x := make(map[int]string)
	x[0] = "hi"
	x[1] = "sundar"
	fmt.Printf("%v", x)
}

func printArray(arr []int) {
	for _, n := range arr {
		fmt.Printf("%d\n", n)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d : %d\n", i, arr[i])
	}
}
