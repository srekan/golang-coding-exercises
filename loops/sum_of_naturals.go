package main
import "fmt"
func getSum(limit int)int{
	sum := 0
	for i := 1; i <= limit; i++{
		sum += i
	}

	return sum
}
func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Printf("Sum of first %d natural numbers = %d\n", n, getSum(n))
}

/** Output

$ go run sum_of_naturals.go
Enter a number: 5
Sum of first 5 natural numbers = 15

$ go run sum_of_naturals.go
Enter a number: 10
Sum of first 10 natural numbers = 55

*/