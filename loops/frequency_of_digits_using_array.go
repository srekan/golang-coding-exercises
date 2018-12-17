package main
import "fmt"
func getFrequencyOfDigits(num int)[10]int{
	// Ten positions for 10 digits -> 0 to 9
	frequencies := [10]int{}

	for i := num; i > 0; i /= 10{
		digit := i % 10
		frequencies[digit]++
	}

	/*
	// Deleting the zero valued keys
	for i := 0; i <= 9; i++{
		if(frequencies[i] == 0){
			delete(frequencies, i)
		}
	}
	*/
	return frequencies
}

func main(){
	fmt.Println("-----Frequency of Digits-----")
	fmt.Println("                 0 1 2 3 4 5 6 7 8 9")
	fmt.Println("                 ...................")
	fmt.Println("1001001     -->", getFrequencyOfDigits(1001001))
	fmt.Println("10221033111 -->", getFrequencyOfDigits(10221033111))
}

/** Output

$ go run frequency_of_digits_using_array.go
-----Frequency of Digits-----
                 0 1 2 3 4 5 6 7 8 9
                 ...................
1001001     --> [4 3 0 0 0 0 0 0 0 0]
10221033111 --> [2 5 2 2 0 0 0 0 0 0]

*/