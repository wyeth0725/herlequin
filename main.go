package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var a = make([]int,N) 
	for i := 0; i < N; i++{
		fmt.Scan(&a[i])
	}
	for i := 0; i < N; i++{
		if a[i] % 2 == 1{
			fmt.Println("first")
			return
		}
	}
	fmt.Println("second")

}
