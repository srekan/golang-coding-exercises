package main
import "fmt"
func main(){
  fmt.Println("-----Arithmetic Operations-----")
  var a = 25
  var b = 2
  fmt.Printf("%d + %d = %d\n", a, b, a + b)
  fmt.Printf("%d - %d = %d\n", a, b, a - b)
  fmt.Printf("%d * %d = %d\n", a, b, a * b)
  fmt.Printf("%d / %d = %.2f\n", a, b, (float64(a) / float64(b)))
}

/** Output - arithmetic_operations.go

-----Arithmetic Operations-----
25 + 2 = 27
25 - 2 = 23
25 * 2 = 50
25 / 2 = 12.50
25 % 2 = 1
  fmt.Printf("%d %% %d = %d\n", a, b, a % b)
}

*/
