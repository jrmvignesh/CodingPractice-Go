
package main

import "testing"

func Testfibonacci(t *testing.T) {
var v int
 v = fibonacci(6)
 if v != 5 {
 t.Error("Expected 5 , got " , v)
 }
 
}
