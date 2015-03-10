package main

import (
  "os/exec"
  "fmt"
  "os"
  "math/cmplx"
  "math"
)

var	TLim = 0.5
var  	TLrl = 0.0
var 	BRim = -1.0
var    	BRrl = 0.5

var     fil *os.File 

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Getc(x, y, w, h int) complex128 {

	xlen := math.Abs(TLrl) + math.Abs(BRrl)
	ylen := math.Abs(TLim) + math.Abs(BRim)

	xpl := (float64(x)/float64(w)) * xlen
	ypl := (float64(y)/float64(h)) * ylen

	xplot := TLrl + xpl
	yplot := BRim + ypl

	fmt.Fprintln(fil, "x y screen cords:", x, y); 
	fmt.Fprintln(fil, "Map xpl ypl on area:", xpl, ypl) 
	fmt.Fprintln(fil, "Map to Mset: real:", xplot,"Img :", yplot);

	
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
		fmt.Fprintln(fil, "Current z:", z, " Radius: ", r, " Angle: ",theta)
		iter++
	}
	_ = theta 
	return iter
}		

func main() {
	var grade string = " .:o*x?OX&#"

	var h, w, iter, maxiter, frac int
	var c complex128
	var err error
 
	fil, err = os.Create("Mand.dat")
	check(err)
		
	defer fil.Close()
	maxiter = 100

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
			iter = Iterate(c, 4.0, maxiter)
			
			frac = int(float64(iter)/float64(maxiter) * 10.0)
			symb := string([]rune(grade)[frac])
			fmt.Print(symb)
		}
		fmt.Print("\n")
	}
}
