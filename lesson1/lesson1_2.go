package main

import (
	"fmt"
)

func main(){
		i := 1
		t,f := true,false
		f64 := 1.2
	fmt.Printf("%T\n",i)
	fmt.Printf("%T\n",t,f)
	fmt.Printf("%T\n",f64)
	fmt.Println(i,t,f,f64)
}

/* $ go run lesson1_2.go 
 int
 bool
 %!(EXTRA bool=false)float64
 1 true false 1.2
*/