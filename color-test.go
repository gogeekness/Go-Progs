package main
 
import "fmt"
 
func main() {
	fmt.Println("\033[1;31m for red text \033[0m")
	for i := 0; i < 8; i++ {

		fmt.Println("\033[",i,";3HHello")
	}

	
}
