package main
import "fmt"
func printNumbersInReverse(n int){
	for i := n; i >= 1; i-- {
		fmt.Println(i)
	}
}
func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	printNumbersInReverse(n)
}

/** Output

$ go run natural_numbers_in_reverse.go
Enter a number: 10
10
9
8
7
6
5
4
3
2
1

*/