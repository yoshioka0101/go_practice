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
