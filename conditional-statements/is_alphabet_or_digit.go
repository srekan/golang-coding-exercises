/*
ASCII codes
0 : 48
9 : 57
A : 65
B : 66
Z : 90
a : 97
b : 98
z : 122  
*/
package main
import (
	"fmt"
	"bufio"
	"os"
)
func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a character: ")
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
		fmt.Println("Given character is an alphabet")
	} else if char >= '0' && char <= '9' {
		fmt.Println("Given character is a digit")
	}
	/*
	// Method 2: compare with ascii values
	if (char >= 65 && char <= 122) || (char >= 97 && char <= 90) {
		fmt.Println("Given character is an alphabet")
	} else if char >= 48 && char <= 57 {
		fmt.Println("Given character is a digit")
	}
	*/
}

/** Outputs

$ go run is_alphabet_or_digit.go
Enter a character: a
Given character is an alphabet

$ go run is_alphabet_or_digit.go
Enter a character: G
Given character is an alphabet

$ go run is_alphabet_or_digit.go
Enter a character: 5
Given character is a digit

*/