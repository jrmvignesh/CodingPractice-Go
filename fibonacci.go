package main

import "fmt"


func main() {

fmt.Print("Enter a number: ")
var x int
var z int
fmt.Scan(&x)
  for i :=1 ; i < x ; i ++ {
  z = fibonacci(i)
  }
 fmt.Println(z)
}

func fibonacci(y int ) int{
 if y <= 1 {
	return y
 }
 
 return fibonacci(y - 1) + fibonacci(y - 2)
 
}
