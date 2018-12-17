package main
import "fmt"
func printTable(n int){
	for i := 1; i <=10; i++{
		fmt.Printf("%d x %d = %d\n", n, i, (n * i))
	}
}
func main(){
	printTable(2)
	fmt.Println("...............")
	printTable(5)
}

/** Output

$ go run multiplication_table.go
2 x 1 = 2
2 x 2 = 4
2 x 3 = 6
2 x 4 = 8
2 x 5 = 10
2 x 6 = 12
2 x 7 = 14
2 x 8 = 16
2 x 9 = 18
2 x 10 = 20
...............
5 x 1 = 5
5 x 2 = 10
5 x 3 = 15
5 x 4 = 20
5 x 5 = 25
5 x 6 = 30
5 x 7 = 35
5 x 8 = 40
5 x 9 = 45
5 x 10 = 50

*/