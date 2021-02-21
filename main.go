package main

import "fmt"

//example-3
func add(a, b int) ( x int){
//body

   x = a+b
   return x

} 

func main() {

   y := add(10, 15)

fmt.Println(y)
}




