package main
import "fmt"
func getFrequencyOfDigits(num int)map[int]int{
	// Ten positions for 10 digits -> 0 to 9
	frequencies := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
	}

	for i := num; i > 0; i /= 10{
		digit := i % 10
		frequencies[digit]++
	}

	// Deleting the zero valued keys
	for i := 0; i <= 9; i++{
		if(frequencies[i] == 0){
			delete(frequencies, i)
		}
	}
	return frequencies
}

func main(){
	fmt.Println("-----Frequency of Digits-----")
	fmt.Println("1001001     -->", getFrequencyOfDigits(1001001))
	fmt.Println("10221033111 -->", getFrequencyOfDigits(10221033111))
}

/** Output

$ go run frequency_of_digits.go
-----Frequency of Digits-----
1001001     --> map[0:4 1:3]
10221033111 --> map[1:5 0:2 2:2 3:2]

*/