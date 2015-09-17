package main

import (
    "fmt"
    "time"
  )

func main() {
c1 := make(chan string)


go func() {
for {
c1 <- "Make cake"
time.Sleep(time.Second * 1)

}
}()

go func() {
for {
msg := <- c1 

fmt.Println(msg)
time.Sleep(time.Second * 2)
fmt.Println("Pack cake")
time.Sleep(time.Second * 2)
time.After(time.Second * 10)
fmt.Println("Delivered")
}
}()

var input string
fmt.Scanln(&input)
}