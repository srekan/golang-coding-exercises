package main
import "fmt"
func main(){
  fmt.Println("-----Simple Interest-----")
  var principle float64
  var numberOFYears float64
  var rateOfInterest float64

  fmt.Print("Enter principle: ")
  fmt.Scan(&principle)
  fmt.Print("Enter number of years: ")
  fmt.Scan(&numberOFYears)
  fmt.Print("Enter rate of interest: ")
  fmt.Scan(&rateOfInterest)

  simpleInterest := principle * numberOFYears * rateOfInterest / 100

  fmt.Printf("Simple Interest = %.2f\n", simpleInterest)
}

/** Output

$ go run simple_interest.go
-----Simple Interest-----
Enter principle: 2000
Enter number of years: 5
Enter rate of interest: 10
Simple Interest = 1000.00

*/