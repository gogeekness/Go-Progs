package main

import "fmt"

func main() {
p := []byte("Hello World")
fmt.Println("p ==", string(p))
fmt.Println("p[1:4] ==", string(p[1:4]))

// missing low index implies 0
fmt.Println("p[:7] ==", string(p[:7]))

// missing high index implies len(s)
fmt.Println("p[4:] ==", string(p[4:]))

r := []byte("####") 
fmt.Println("r[1:4] ==", string(r[1:4]))


copy(p[:3], r)
fmt.Println("p[:] ==", string(p))




}
