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
func main(){
	fmt.Println("-----Is Prime-----")
	fmt.Println("2   --->", isPrime(2))
	fmt.Println("11  --->", isPrime(11))
	fmt.Println("12  --->", isPrime(12))
	fmt.Println("13  --->", isPrime(13))
	fmt.Println("14  --->", isPrime(14))
}

/** Output

$ go run is_prime.go
-----Is Prime-----
2   ---> true
11  ---> true
12  ---> false
13  ---> true
14  ---> false

*/