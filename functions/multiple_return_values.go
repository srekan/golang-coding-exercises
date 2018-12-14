package main
import "fmt"

func getSumAndProduct(a, b int)(int, int){
	return (a+b), (a*b)
}

func main(){
	sum, product := getSumAndProduct(20, 2)
	fmt.Printf("20 + 2 = %d\n", sum)
	fmt.Printf("20 * 2 = %d\n", product)

	fmt.Println()
	sum, product = getSumAndProduct(30, 3)
	fmt.Printf("30 + 3 = %d\n", sum)
	fmt.Printf("30 * 3 = %d\n", product)
}

/** Output

$ go run multiple_return_values.go
20 + 2 = 22
20 * 2 = 40

30 + 3 = 33
30 * 3 = 90

*/