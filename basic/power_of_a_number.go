package main
import "fmt"
import "math"
func main(){
  fmt.Println("-----Power of a number: x^y-----")
  var base float64
  var exponent float64

  fmt.Print("Enter base: ")
  fmt.Scan(&base)
  fmt.Print("Enter exponent: ")
  fmt.Scan(&exponent)

  power := math.Pow(base, exponent)
  fmt.Printf("%.0f ^ %.0f = %.0f\n", base, exponent, power)
}

/** Output

$ go run power_of_a_number.go
-----Power of a number: x^y-----
Enter base: 2
Enter exponent: 3
2 ^ 3 = 8

*/