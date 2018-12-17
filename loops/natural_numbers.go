package main
import "fmt"
func printNumbers(n int){
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	printNumbers(n)
}

/** Output

$ go run natural_numbers.go
Enter a number: 10
1
2
3
4
5
6
7
8
9
10

*/