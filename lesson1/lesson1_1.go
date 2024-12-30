package main

import (
	"fmt"
	"os/user"
	"time"
)
func init(){
	fmt.Println("init!!")
}

func buzz(){
	fmt.Println("Bazz")
}

func main(){
	buzz()
	fmt.Println("Hello world!!", time.Now())
	fmt.Println(user.Current())
}