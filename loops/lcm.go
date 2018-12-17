package main
import "fmt"
func getLCM(num1, num2 int)int{
	var max int
	if num1 > num2 {
		max = num1
	} else {
		max = num2
	}
	
    // First multiple to be checked
	i := max
	var lcm int
    
    // Run loop indefinitely till LCM is found
    for {
        if i % num1 == 0 && i % num2 == 0 {
            /*
             * If 'i' divides both 'num1' and 'num2'
             * then 'i' is the LCM.
             */
            lcm = i;

            /* Terminate the loop after LCM is found */
            break
        }

        /*
         * If LCM is not found then generate next 
         * multiple of max between both numbers
         */
        i += max
	}
	
	return lcm
}

func main(){
	fmt.Println("-----LCM-----")
	fmt.Println("LCM of 12 and 30:", getLCM(12, 30))
	fmt.Println("LCM of  4 and 10:", getLCM(4, 10))
}

/** Output

$ go run lcm.go
-----LCM-----
LCM of 12 and 30: 60
LCM of  4 and 10: 20

*/