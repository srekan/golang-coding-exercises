package main
import "fmt"
func main(){
	type Month struct {
		name string
		days int
	}
	monthsMap := map[int]Month{
		1: { name: "January", days: 31 },
		2: { name: "February", days: 28 },
		3: { name: "March", days: 31 },
		4: { name: "April", days: 30 },
		5: { name: "May", days: 31 },
		6: { name: "June", days: 30 },
		7: { name: "July", days: 31 },
		8: { name: "August", days: 31 },
		9: { name: "September", days: 30 },
		10: { name: "October", days: 31 },
		11: { name: "November", days: 30 },
		12: { name: "December", days: 31 },
	}

	var n int
	fmt.Print("Enter month number(1 - 12): ")
	fmt.Scan(&n)

	if monthsMap[n].name != "" {
		fmt.Printf("%s - %d days\n", monthsMap[n].name, monthsMap[n].days)
	} else {
		fmt.Println("Please enter a valid value between 1 and 12")
	}
}

/** Outputs

$ go run days_of_month_using_map.go
Enter month number(1 - 12): 6
June - 30 days

$ go run days_of_month_using_map.go
Enter month number(1 - 12): 66
Please enter a valid value between 1 and 12

*/