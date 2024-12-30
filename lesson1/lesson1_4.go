package main

import "fmt"

func main() {
	n := map[string]int{"apple": 100 , "banana": 200}
	fmt.Println(n)
	// appleのキーと紐づいたバリューを出力する
	fmt.Println(n["apple"])
	// bananaの値を300にする
	n["banana"] = 300
	fmt.Println(n)
	//　Suikaを追加する
	n["Suika"] = 500
	fmt.Println(n)
}
