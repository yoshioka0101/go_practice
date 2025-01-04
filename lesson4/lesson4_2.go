package main

import "fmt"

type Animal struct{}

func (a Animal) Speake() string{
    return "Iam Animal"
}

type Docs struct{
    Animal
}

func main(){
    docs := Docs{}
    fmt.Println(docs.Speake())
}
