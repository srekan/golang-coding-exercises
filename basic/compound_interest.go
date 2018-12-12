/**
 CI = P * (1 + (R / 100))^T
*/
package main
import (
	"math"
	"fmt"
)
func main(){
  fmt.Println("-----Compound Interest-----")
  var principle float64
  var numberOFYears float64
  var rateOfInterest float64

  fmt.Print("Enter principle: ")
  fmt.Scan(&principle)
  fmt.Print("Enter number of years: ")
  fmt.Scan(&numberOFYears)
  fmt.Print("Enter rate of interest: ")
  fmt.Scan(&rateOfInterest)

  compoundInterest := principle * math.Pow((1 + rateOfInterest / 100), numberOFYears)

  fmt.Printf("Compound Interest = %.2f\n", compoundInterest)
}

/** Output

$ go run compound_interest.go
-----Compound Interest-----
Enter principle: 2000
Enter number of years: 5
Enter rate of interest: 10
Compound Interest = 3221.02

*/