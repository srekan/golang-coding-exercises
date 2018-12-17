package main
import "fmt"
func getFactors(n int)[]int{
	var factors []int 
	for i := 1; i <= n; i++ {
        if n % i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
func main(){
	fmt.Println("-----Printing Factors-----")
	fmt.Println("12  -->", getFactors(12))
	fmt.Println("121  -->", getFactors(121))
	fmt.Println("1200  -->", getFactors(1200))
}

/** Output

$ go run find_factors.go
-----Printing Factors-----
12  --> [1 2 3 4 6 12]
121  --> [1 11 121]
1200  --> [1 2 3 4 5 6 8 10 12 15 16 20 24 25 30 40 48 50 60 75 80 100 120 150 200 240 300 400 600 1200]

*/