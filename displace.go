
package main

import ( 
	"fmt"
	"math"
	)

func GenDisplaceFn(a float64, v_o float64, s_o float64) func(t float64) float64	{
	fn:= func (t float64) float64{
		s := 0.5 * a * math.Pow(t, 2) + v_o * t + s_o
		return s
	}
	return fn
}



var a_prompt string = "Enter acceleration: "
var v_prompt string = "Enter initial velocity: "
var d_prompt string = "Enter initial displacement: "
var t_prompt string = "Enter time: "

func main(){
	var accel float64
	var velocity float64
	var i_displace float64
	var time float64

	fmt.Print(a_prompt)
	fmt.Scan(&accel)

	fmt.Print(v_prompt)
	fmt.Scan(&velocity)

	fmt.Print(d_prompt)
	fmt.Scan(&i_displace)

	fmt.Print(t_prompt)
	fmt.Scan(&time)

	fn := GenDisplaceFn(accel, velocity, i_displace)
	fmt.Printf("Displacement after elapsed time: %.2f seconds\n", fn(time))
	
}







