# 1章

## 出力
- 改行を使用するときは
    - `\n`を使用することができる
    - または`Print`ではなく、`Println`をすることで改行を指定できる

```go
package main

import "fmt"

func main(){
	fmt.Print("Hello world\n")
	fmt.Println("Hello world!!")
	fmt.Println("Hello world!!", "Go Projet")
}

/* 出力結果
$ go run lesson.go 
Hello world
Hello world!!
Hello world!! Go Projet
*/
```


```go
package main

import "fmt"

func buzz(){
	fmt.Println("Bazz")
}

func main(){
	buzz()
	fmt.Print("Hello world\n")
	fmt.Println("Hello world!!")
	fmt.Println("Hello world!!", "Go Projet")
}

/*出力結果
$ go run lesson.go 
Bazz
Hello world
Hello world!!
Hello world!! Go Projet
*/
```

- init関数
    - init関数は最初に実行される

```go
package main

import "fmt"

// １：最初に実行される
func init(){
	fmt.Println("init!!")
}

//　３：main関数で呼び出されて実行される
func buzz(){
	fmt.Println("Bazz")
}

// ２：init関数実行後にmain関数を実行される
func main(){
	buzz()
	fmt.Println("Hello world!!", "Go Projet")
}

/*出力結果
$ go run lesson.go 
init!!
Bazz
Hello world!! Go Projet
*/
```

- 他のパッケージをインストールする
  - 今回は`os/user`と`time`を使用する

```go
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

/*出力結果
$ go run lesson.go 
init!!
Bazz
Hello world!! 2024-12-29 19:24:38.224148947 +0900 JST m=+0.000025668
&{501 1000 opm007756 OPM007756 /home/opm007756.linux} <nil>
*/
```


## 変数宣言
- varで変数を宣言することができる
- ちなみに制限するときに`var t,f bool`にしてtとfを`true`と`false`で指定しない場合は両方とも`false`になる
    - 理由はデフォルトが`false`だから
- またvarを複数定義する際には、var()でくくることもできる

```go
package main

import (
	"fmt"
)

func main(){
	var (
		i int = 1
		t,f bool = true,false
		f64 float64 = 1.2
	)
	fmt.Println(i,t,f,f64)
}

/* 出力結果
$ go run lesson1_2.go 
1 true false
*/
```

- `:=`にするとvarを書く必要がなくなる
- もし型を把握したい場合は、`Printf("%T",変数)`にすると型を出力することができる

```go
package main

import (
	"fmt"
)

func main(){
		i := 1
		t,f := true,false
		f64 := 1.2
	fmt.Printf("%T\n",i)
	fmt.Printf("%T\n",t,f)
	fmt.Printf("%T\n",f64)
	fmt.Println(i,t,f,f64)
}

package main

import (
	"fmt"
)

func main(){
		i := 1
		t,f := true,false
		f64 := 1.2
	fmt.Printf("%T\n",i)
	fmt.Printf("%T\n",t,f)
	fmt.Printf("%T\n",f64)
	fmt.Println(i,t,f,f64)
}

/* 出力結果
$ go run lesson1_2.go 
 int
 bool
 %!(EXTRA bool=false)float64
 1 true false 1.2
*/
```

### データ型

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	var s string = "Hello World"
	fmt.Println(s)
	s = strings.Replace(s, "H","X",1)
	fmt.Println(s)
}

/*出力結果
$ go run lesson1_3.go 
Hello World
Xello World
*/
```

### 配列操作
- 配列は以下のように記載する
- また配列を追加する場合に`append`を使用することで要素を追加することができる

```go
package main

import "fmt"

func main() {
	n := []int{1,2,3,4,5}
	fmt.Println(n)
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	n[2] = 1000
	fmt.Println(n)
	n = append(n,100,200,300)
	fmt.Println(n)
}

/*出力
$ go run lesson1_4.go 
[1 2 3 4 5]
[3 4]
[1 2]
[3 4 5]
[1 2 1000 4 5]
[1 2 1000 4 5 100 200 300]
*/
```

#### 多次元配列

```go
package main

import "fmt"

func main() {
	var board = [][]int{
		[]int{0, 1 ,2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)
	fmt.Println(board[1])
	fmt.Println(board[1][2])
	fmt.Println(board[2][2])
	fmt.Println(board[0][0])
}

/*出力
$ go run lesson1_4.go 
[[0 1 2] [3 4 5] [6 7 8]]
[3 4 5]
5
8
0
*/

```

- make関数

```go
package main

import "fmt"

func main() {
	n := make([]int, 3, 5)
	fmt.Println(n)
	n = append(n,0,0)
	fmt.Println(n)
	n = append(n,1,2,3,4,5,6)
	fmt.Println(n)
}
/*出力
$ go run lesson1_4.go 
[0 0 0]
[0 0 0 0 0]
[0 0 0 0 0 1 2 3 4 5 6]
*/
```

- `map`関数を用いた多次元配列


```go
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

// $ go run lesson1_4.go 
// map[apple:100 banana:200]
// 100
// map[apple:100 banana:300]
// map[Suika:500 apple:100 banana:300]

```