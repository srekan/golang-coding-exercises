package main
import "fmt"
func main(){
  fmt.Println("----Finding area of a triangle-----")
  var height float32
  var breadth float32

  fmt.Print("Enter height: ")
  fmt.Scan(&height)
  
  fmt.Print("Enter base: ")
  fmt.Scan(&breadth)

  area := height * breadth
  fmt.Println("Area =", area)
}

/** Output

$ go run area_of_triangle.go
----Finding area of a triangle-----
Enter height: 5.6
Enter base: 6.7
Area = 37.519997

*/
