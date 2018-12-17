package main
import "fmt"
func printEvens(limit int){
	for i := 0; i <= limit; i += 2 {
		fmt.Println(i)
	}
}
func main(){
	printEvens(100)
}
