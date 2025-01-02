package	main

import "fmt"

func main() {
    // ポインタを作成する
    p := new(int)

    // ポインタのアドレスを出力する
    fmt.Println(p) // => 0xc0000a6008

    // ポインタの値を設定する
    *p = 123

    // ポインタの値を出力する
    fmt.Println(*p) // => 123
}
