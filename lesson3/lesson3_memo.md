# 3章
## ポインタ

1. 普通の変数に値を格納
  - 最初に通常の変数を宣言して値を設定

2. その変数のアドレスをポインタに保存
  - `&`演算子を使って変数のアドレスを取得し、ポインタ型の変数に保存する

3. 再度値を扱う際にポインタを使用
  - ポインタを使って値を参照または変更することで、変数そのものを直接操作せずにメモリ効率を向上させる

```go
package main

import "fmt"

func main() {
	// 1. 普通の変数に値を格納
	var n int = 100
	fmt.Println("Step 1: 変数 n の値:", n)

	// 2. ポインタに変数のアドレスを保存
	var ptr *int = &n
	fmt.Println("Step 2: ポインタ ptr の値（n のアドレス）:", ptr)

	// 3. ポインタを使用して値を参照
	fmt.Println("Step 3: ポインタを使って参照した値:", *ptr)

	// 4. ポインタを使って値を変更
	*ptr = 200
	fmt.Println("Step 4: ポインタを使って変更後の n の値:", n)

	// 5. メモリ効率の確認（変数ではなくポインタを通じて操作）
	fmt.Printf("n のアドレス: %p, ポインタ ptr のアドレス: %p\n", &n, ptr)
}

```

```sh
$ go run lesson3_1.go 

Step 1: 変数 n の値: 100
Step 2: ポインタ ptr の値（n のアドレス）: 0x400000e0d0
Step 3: ポインタを使って参照した値: 100
Step 4: ポインタを使って変更後の n の値: 200
n のアドレス: 0x400000e0d0, ポインタ ptr のアドレス: 0x400000e0d0

```

### Struct

- 構造体
  - Railsにおけるインスタンス変数に近い

```go
package	main

import "fmt"

type User struct {
	ID int
	Name string
}

func main() {
	user := User{
		ID: 1,
		Name: "Shishi",
	}

	fmt.Println("ID:", user.ID)
	fmt.Println("Name:", user.Name)
}

```