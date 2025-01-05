package main

import "fmt"

func do(i interface{}){
    ss := i.(string)
    fmt.Println(ss + "!")
}

func main(){
    do("Mike")
    do(123)
}

/*

do("Mike)はstringで定義されているので成功するが、do(123)は定義されていないのでpanicとして扱われます
$ go run lesson4_4.go 
Mike!
panic: interface conversion: interface {} is int, not string

goroutine 1 [running]:
main.do({0xa8920?, 0xe47e0?})
        /home/opm007756.linux/go_book/go_professional/lesson4/lesson4_4.go:6 +0xa4
main.main()
        /home/opm007756.linux/go_book/go_professional/lesson4/lesson4_4.go:12 +0x40
exit status 2
*/