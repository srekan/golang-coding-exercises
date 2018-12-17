package main
import "fmt"
func isPrime(num int)bool{
	var hasFactors = true
	for i := 2; i <= num / 2; i++ {
		if(num % i == 0){
			hasFactors = false
			break
		}
	}

	return hasFactors
}
func getPrimeFactors(num int)[]int{
	primeFactors := []int{}
	for i := 2; i <= num / 2; i++{
		if num % i == 0 && isPrime(i){
			primeFactors = append(primeFactors, i)
		}
	}

	return primeFactors
}
func main(){
	fmt.Println("-----Printing Prime Factors-----")
	fmt.Println("Prime Factors of 12 :", getPrimeFactors(12))
	fmt.Println("Prime Factors of 100:", getPrimeFactors(100))
	fmt.Println("Prime Factors of 121:", getPrimeFactors(121))
}

/** Output

$ go run prime_factors.go
-----Printing Prime Factors-----
Prime Factors of 12 : [2 3]
Prime Factors of 100: [2 5]
Prime Factors of 121: [11]

*/