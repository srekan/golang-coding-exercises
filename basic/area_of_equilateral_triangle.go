package main
import "fmt"
import "math"
func main(){
  fmt.Println("-----Area of equilateral triangle-----")
  var length float64
  fmt.Print("Enter length of a side (in meters): ")
  fmt.Scan(&length)

  area := (math.Sqrt(3) / 4) * length * length
  fmt.Printf("Area = %.2f square meters\n", area)
}

/** Output

$ go run area_of_equilateral_triangle.go
-----Area of equilateral triangle-----
Enter length of a side (in meters): 5
Area = 10.83 square meters

*/