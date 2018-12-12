package main
import "fmt"
func main(){
  fmt.Println("-----Convert Days to Years and Months-----")
  var days int
  fmt.Print("Enter number of days: ")
  fmt.Scan(&days)

  years := days / 365
  weeks := (days % 365) / 7
  days  = days - ((years * 365) + (weeks * 7))

  fmt.Printf("Years: %d\n", years);
  fmt.Printf("Weeks: %d\n", weeks);
  fmt.Printf("Days : %d\n", days);
}

/** Output
$ go run days_to_years.go
-----Convert Days to Years and Months-----
Enter number of days: 417
Years: 1
Weeks: 7
Days : 3

*/