package main
import "fmt"
func main(){
	var n int
	fmt.Print("Enter week number(1 - 7): ")
	fmt.Scan(&n)

	if n == 1{
		fmt.Println("Monday")
	} else if n == 2{
		fmt.Println("Tuesday")
	} else if n == 3{
		fmt.Println("Wednesday")
	} else if n == 4{
		fmt.Println("Thursday")
	} else if n == 5 {
		fmt.Println("Friday")
	} else if n == 6{
		fmt.Println("Saturday")
	} else if n == 7{
		fmt.Println("Sunday")
	} else{
		fmt.Println("Please enter a valid value between 1 and 7")
	}
}

/** Outputs

$ go run print_week_day.go
Enter week number(1 - 7): 4
Thursday

$ go run print_week_day.go
Enter week number(1 - 7): 7
Sunday

$ go run print_week_day.go
Enter week number(1 - 7): 8
Please enter a valid value: 1 - 7

*/