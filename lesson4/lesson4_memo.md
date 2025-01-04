# 4章
## メソッド
- 以下のコードでは、

- 通常の関数
  - `func Area(v Vertex) int`は、`Vertex`構造体のインスタンスを引数として受け取る通常の関数
- メソッド
  - `func (v Vertex) Area() int`は、`Vertex`構造体に紐づいたメソッド
  - メソッドは、構造体（レシーバー）に直接関連付けられ、呼び出す際には`v.Area()`のようにドット記法を使用する

```go
package main

import "fmt"

type Vertex struct{
	X,Y int
}

func Area(v Vertex) int{
	return v.X * v.Y
}

func (v Vertex) Area() int{
	return v.X * v.Y
}

func main(){
	v:= Vertex{3,4}
	fmt.Println(Area(v))
	fmt.Println(v.Area())
}
```


#### ポイントレシーバーと値レシーバの違いを説明している
- 値レシーバーはレシーバーがコピーされるため、メソッド内での影響を与えない
- それに対してポイントレシーバーは元のレシーバーの値を書き換えることができる

そのため、レシーバーを書き換えたいかどうかでどちらを使用するか考えるべき
多分Web開発いなどではポインターレシーバーgがよく利用されそう


```go
package main

import "fmt"

type Vertex struct{
	X,Y int
}

func (v Vertex) Area() int{
	return v.X * v.Y
}

func (v *Vertex) Scale(i int){
	v.X = v.X * i
	v.Y = v.Y * i
}

func main(){
	v:= Vertex{3,4}
	fmt.Println("通常の関数:",v.Area())
	v.Scale(10)
	fmt.Println("通常の関数:",v.Area())
}
```

####  継承

- 資料にある問題ではなくて、Rubyでよくある親クラス、子クラスの話をしてみた
  - AnimalはいかにDocが存在していて、Animalで定義したSpeakメソッドをDocで使用する


```go
package main

import "fmt"

type Animal struct{}

func (a Animal) Speake() string{
    return "Iam Animal"
}

// ここで継承元のメソッドを使用できる
type Docs struct{
    Animal
}

func main(){
    // Rails のnew.ckass名みたいなもの
    doc := Doc{}
    fmt.Println(doc.Speak)
}

```