package main

import "fmt"

func main() {
        fmt.Println("\033[1;31m for red text \033[0m")
        for i := 9600; i < 9680; i++ {
		if (i % 50) == 0 {
			fmt.Println()
		}
                fmt.Printf("%s", string(i))
        }
	fmt.Println()

}
