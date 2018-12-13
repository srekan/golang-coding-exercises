package main
import (
	"fmt"
	"bufio"
	"os"
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a letter from alphabet: ")
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	if char >= 'a' && char <= 'z' {
		fmt.Println("LOWER_CASE")
	} else if char >= 'A' && char <= 'Z' {
		fmt.Println("UPPER_CASE")
	} else {
		fmt.Println("NOT_AN_ALPHABET")
	}
}

/** Outputs

$ go run is_upper_case.go
Enter a letter from alphabet: g
LOWER_CASE

$ go run is_upper_case.go
Enter a letter from alphabet: G
UPPER_CASE

$ go run is_upper_case.go
Enter a letter from alphabet: 8
NOT_AN_ALPHABET

*/