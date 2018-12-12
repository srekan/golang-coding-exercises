package main
import "fmt"
func main(){
  fmt.Println("-----Convert Celcius to Fahrenheit-----")
  var celcius float64
  fmt.Print("Enter temperature in celcius degrees: ") 
  fmt.Scan(&celcius)
  fahrenheit  := (celcius * (9/5)) + 32
  fmt.Println("Temperature in Fahrenheit degrees= ", fahrenheit)
}

/** Output

$ go run celsius_to_fahrenheit.go
-----Convert Celcius to Fahrenheit-----
Enter temperature in celcius degrees: 30
Temperature in Fahrenheit degrees=  86

*/
