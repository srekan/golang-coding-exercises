package main
import "fmt"
func getSum(limit int)int{
	sum := 0
	for i := 0; i <= 2 * limit; i+=2{
		sum += i
	}

	return sum
}
func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Printf("Sum of first %d even numbers = %d\n", n, getSum(n))
}

/** Output

$ go run sum_of_evens.go
Enter a number: 5
Sum of first 5 even numbers = 30

$ go run sum_of_evens.go
Enter a number: 10
Sum of first 10 even numbers = 110

*/