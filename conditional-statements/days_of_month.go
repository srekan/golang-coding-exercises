package main
import "fmt"
func main(){
	var n int
	fmt.Print("Enter month number(1 - 12): ")
	fmt.Scan(&n)

	switch n {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 days")
	case 2:
		fmt.Println("28/29 days")
	case 4, 6, 9, 11:
		fmt.Println("30 days")
	default:
		fmt.Println("Please enter a valid value between 1 and 12")
	}
}

/** Outputs

$ go run days_of_month.go
Enter month number(1 - 12): 5
31 days

$ go run days_of_month.go
Enter month number(1 - 12): 2
28/29 days

$ go run days_of_month.go 13
Enter month number(1 - 12): 13
Please enter a valid value between 1 and 12

*/