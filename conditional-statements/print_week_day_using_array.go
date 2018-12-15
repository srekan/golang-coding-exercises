package main
import "fmt"
func main(){
	var n int
	fmt.Print("Enter week number(1 - 7): ")
	fmt.Scan(&n)
	/* Note: There is a dedicated chapter to demonstrate arrays.
	 * This purpose of this program is to show the easy method to 'print_week_day_using_switch.go'
	 */
	weekDays := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
 
	if(n >= 1 && n <= 7){
		fmt.Println(weekDays[n - 1]) // Array index starts from 0.
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