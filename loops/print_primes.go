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
	fmt.Println("-----Printing Primes till 100-----")
	for i := 2; i <= 100; i++{
		if(isPrime(i)){
			fmt.Println(i)
		}
	}
}

/** Output

$ go run print_primes.go
-----Printing Primes till 100-----
2
3
5
7
11
13
17
19
23
29
31
37
41
43
47
53
59
61
67
71
73
79
83
89
97

*/