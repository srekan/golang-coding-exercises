package main
import "fmt"
func main(){
  fmt.Println("-----Input from User-----")
  var myVal int
  fmt.Print("Enter a number: ")
  fmt.Scan(&myVal)
  fmt.Printf("Thank you! You entered %d!\n", myVal)
}

/** Output

$ go run input_from_user.go
-----Input from User-----
Enter a number: 234
Thank you! You entered 234!

*/
