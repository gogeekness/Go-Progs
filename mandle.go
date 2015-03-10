// Mandle set 1
//
// Richard Eseke
package main

import (
	"fmt"
	//"math"
)

func sqrz(z, c complex64) complex64 {
	return (z * z) + c
}

func limit(z complex64, l float64, nt int) complex64 {
	C := 0 
	for i := 1; i <= nt; i++ {
		if (z * z) < l {
			z = sqrz(z, c)
		}	
			
	}
	return z
}	

func main() {
	fmt.Println(sqrz(3+2i, 5+9i))

}	
  
