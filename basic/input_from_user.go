package main
import "fmt"
func main(){
  fmt.Println("-----Input from User-----")
  var my_val int
  fmt.Print("Enter a number: ")
  fmt.Scan(&my_val)
  fmt.Printf("Thank you! You entered %d!\n", my_val)
}

/** Output input_from_user.go

-----Input from User-----
Enter a number: 234
Thank you! You entered 234!

*/
