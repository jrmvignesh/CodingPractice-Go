
package main

import "testing"



func TestShape(t *testing.T ) {

    r := rect{width: 20, height: 4}
    c := circle{radius: 10}


    measure(r)
    measure(c)
    
}
