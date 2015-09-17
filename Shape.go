
package main

import "fmt"
import "math"


type Shape interface {
    Area() float64
    Perimeter() float64
}


type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}


func (r rect) Area() float64 {
    return r.width * r.height
}
func (r rect) Perimeter () float64 {
    return 2*r.width + 2*r.height
}


func (c circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) Perimeter() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g Shape) {
   // fmt.Println(g)
    fmt.Println(g.Area())
    fmt.Println(g.Perimeter())
}

func main() {
    //r := rect{width: 5, height: 4}
    //c := circle{radius: 7}
fmt.Print("Enter dimensions for rectangle: ")
var x,z float64
fmt.Scan(&x)
fmt.Scan(&z)
    r := rect{width: x, height: z}
  
measure(r)
fmt.Print("Enter dimensions for circle: ")
var y float64
fmt.Scan(&y)
  c := circle{radius: y}
measure(c)
}
