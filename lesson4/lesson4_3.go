package main

import "fmt"

// 1: Humanインターフェースを定義
// メソッドの宣言だけで処理内容は記載しない。
type Human interface {
    Say()
}

// 2: Person構造体を定義
// Nameフィールドを持つ構造体。
type Person struct {
    Name string
}

// 3: Person構造体にSayメソッドを実装
// この実装により、Person型はHumanインターフェースを満たす。
func (p Person) Say() {
    fmt.Println(p.Name)
}

// 4: main関数の開始
func main() {
    // 実行のエントリポイント

    // 5: Person型のインスタンスを作成し、Human型に代入
    // インターフェース型のmikeにPerson{"Mike"}を代入。
    var mike Human = Person{"Mike"}

    // 6: インターフェースHumanのSayメソッドを呼び出し
    // 実際には、Person型のSayメソッドが実行される。
    mike.Say()
}
