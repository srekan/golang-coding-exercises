package main
import "fmt"

/* Function definition */
func printMessage(){
	fmt.Println("My Name is Srekan. I am a GO developer. I love Go.")
}

func main(){
	// We will invoke the function three times
	printMessage() // Invoking function a.k.a calling a function
	printMessage()
	printMessage()
}

/** Output

$ go run simple_function.go
My Name is Srekan. I am a GO developer. I love Go.
My Name is Srekan. I am a GO developer. I love Go.
My Name is Srekan. I am a GO developer. I love Go.

*/