package main

import "fmt"

func main() {
	fmt.Println("Vòng lặp for nè")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("Vòng lặp while nè")
    sum_while := 1
    	for sum_while < 1000 {
    		sum_while += sum_while
    	}
    	fmt.Println(sum_while)
}