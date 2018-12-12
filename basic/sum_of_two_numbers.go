package main
import "fmt"
func main(){
  fmt.Println("----Sum of two numbers-----")
  var a = 10
  var b = 20
  var sum = a + b
  fmt.Println("a = ", a)
  fmt.Println("b = ", b)
  fmt.Println(a, "+", b,  "=", sum)
}

/** Outputs

$ go run sum_of_two_numbers.go
----Sum of two numbers-----
a =  10
b =  20
10 + 20 = 30

*/ 
