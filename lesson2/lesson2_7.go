package main

import "fmt"

func thirdPartyConnectDB(){
	panic("Unable to connect database")
}

func save(){
	thirdPartyConnectDB()

	defer func() {
		s := recover()
		fmt.Println(s)
	}()
}

func main(){
	save()
	fmt.Println("OK?")
}