package main
import "fmt"
func getSum(limit int)int{
	sum := 0
	for i := 1; i <= 2 * limit; i+=2{
		sum += i
	}

	return sum
}
func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Printf("Sum of first %d odd numbers = %d\n", n, getSum(n))
}

/** Output

$ go run sum_of_odds.go
Enter a number: 5
Sum of first 5 odd numbers = 25

$ go run sum_of_odds.go
10Enter a number:
Sum of first 10 odd numbers = 100

*/