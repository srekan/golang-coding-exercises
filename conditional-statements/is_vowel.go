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
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
		switch char {
			case 'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u':
				fmt.Println("VOWEL")
			default:
				fmt.Println("CONSONANT")
		}
	} else {
		fmt.Println("The input is not an alphabet")
	}
}

/**

$ go run is_vowel.go
Enter a letter from alphabet: A
VOWEL

$ go run is_vowel.go
Enter a letter from alphabet: g
CONSONANT

$ go run is_vowel.go
Enter a letter from alphabet:6
The input is not an alphabet

*/