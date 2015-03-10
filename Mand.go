package main

import (
  "os/exec"
  "fmt"
  "os"
  "math/cmplx"
  "math"
)

var	TLim = -0.405
var  	TLrl = 0.335
var 	BRim = -0.410
var    	BRrl = 0.340

func Getc(x, y, w, h int) complex128 {

	xlen := math.Abs(TLrl - BRrl)
	ylen := math.Abs(TLim - BRim)

	xpl := (float64(x)/float64(w)) * xlen
	ypl := (float64(y)/float64(h)) * ylen

	xplot := TLrl + xpl
	yplot := BRim + ypl

	c := complex(xplot, yplot)
	return c 
}


func Iterate(c complex128, limit float64, num int) int {
	z := 0+0i
	ztmp := 0+0i
	r := 0.0
	var theta float64
	iter := 1

	for iter < num && r < limit{
		ztmp = cmplx.Pow(z, 2) + c
		r, theta = cmplx.Polar(z)
		z = ztmp
		iter++
	}
	_ = theta 
	return iter
}		

//func ColorScale() string {
//	str string
//	for i:=1; i < 8; i++ {
		

func main() {
	var prnt string

	var h, w, iter int
	
	var c complex128

	var maxiter float64 = 10000.0

	// ## Get size of terminal ##
	cmd:= exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	d, _ := cmd.Output()
	fmt.Sscan(string(d), &h, &w)
 
        // ## Clear the srceen ##
        cmd2:= exec.Command("clear")
        cmd2.Stdout = os.Stdout
        cmd2.Run()

	for i:= 1; i < h; i++ {
		for j:= 1; j < w; j++ {
			c = Getc(j, i, w, h)
			iter = Iterate(c, 4.0, int(maxiter))
			frac := int( math.Log( float64(iter) ) )
			if frac > 7 {
				frac = 7 }
			if iter >= int(maxiter) {
				frac = 0 }
		        prnt = fmt.Sprintf("\x1b[%dm%s\x1b[0m", 30+frac, string(9608))
			fmt.Printf("%s", prnt)

		}
		fmt.Printf("\n")
	}

}
