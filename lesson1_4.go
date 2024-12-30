package main

import "fmt"

func main() {
	n := []int{1,2,3,4,5}
	fmt.Println(n)
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	n[2] = 1000
	fmt.Println(n)
	n = append(n,100,200,300)
	fmt.Println(n)
}
