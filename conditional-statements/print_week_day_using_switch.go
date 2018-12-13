package main
import "fmt"
func main(){
	var n int
	fmt.Print("Enter week number(1 - 7): ")
	fmt.Scan(&n)

	switch n {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
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