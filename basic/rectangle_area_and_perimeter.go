package main
import "fmt"
func main(){
  fmt.Println("-----Find Reactangle Area and Perimeter-----")
  var length float64
  var breadth float64

  fmt.Print("Enter length of rectangle (in meters): ")
  fmt.Scan(&length)
  fmt.Print("Enter breadth of rectangle (in meters): ")
  fmt.Scan(&breadth)

  area := length * breadth
  perimeter := 2 * (length + breadth)
  
  fmt.Println("Length   :", length, "meters")
  fmt.Println("Breadth  :", breadth, "meters")
  fmt.Println("Area     :", area, "square meters")
  fmt.Println("Perimeter:", perimeter, "meters")
}

/** Output

$ go run rectangle_area_and_perimeter.go
-----Find Reactangle Area and Perimeter-----
Enter length of rectangle (in meters): 2.5
Enter breadth of rectangle (in meters): 5.5
Length   : 2.5 meters
Breadth  : 5.5 meters
Area     : 13.75 square meters
Perimeter: 16 meters

*/
