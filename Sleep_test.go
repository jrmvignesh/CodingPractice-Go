package main

import (
    "fmt"
    "time"
    "testing"
  )
  
  func TestSleep(t *testing.T ) {

c := make(chan string)


go func() {
for {
c <- "Testing"
time.Sleep(time.Second * 1)
}
}()



go func() {
for {
select {
case msg1 := <- c:
fmt.Println(msg1)
fmt.Println(time.Now())

case  <- time.After(time.Second):
fmt.Println("timeOut")
}
}
}()
var input string
fmt.Scanln(&input)
}