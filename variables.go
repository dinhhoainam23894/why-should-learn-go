package main

import "fmt"

func main() {
   var x float64 = 20.0
   y := 42
   var a, b, c = 3, 4, "foo"
   fmt.Println(x , y , a , b , c)
   fmt.Printf("x is of type %T\n", x)
   fmt.Printf("y is of type %T\n", y)
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n", c)
}