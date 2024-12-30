# 2章
## 条件分岐

- if文は一般的な表現だと思う

```go
package	main

import "fmt"

func main(){
	num := 1
	if num % 2 == 0{
		fmt.Println("by 2")
	}else{
		fmt.Println("not 2")
	}
}
```

### for文

- for文も基本そこまで特殊でもない
- continueを入れる
  - 良さわかってない

```go
package	main

import "fmt"

func main(){
	for i :=0; i < 10; i++{
		if i == 4 {
			fmt.Println("continueを入れる")
			continue
		}
		fmt.Println(i)
	}

}
```

- `range`文

- 普通にforで書く内容をrangeにしていく実装を書いていく

```go
package	main

import "fmt"

func main(){
	l := []string{"Go","Ruby","PHP"}

	for i := 0; i < len(l); i++{
		fmt.Println(i, l[i])
	}

}
```

- `range`は、スライス、配列、マップなどを反復処理するための構文
- `range`を使うと、スライスや配列のインデックスと値を順番に取得できる


```go
package	main

import "fmt"

func main(){
	l := []string{"Go","Ruby","PHP"}

	for _, v := range l{
		fmt.Println(v)
	}

}
```

- key,valueを出力する

```go
package	main

import "fmt"

func main(){
	m := map[string]int{"Apple": 100, "Orange": 200, "Tomato": 300}

	for k, v := range m{
		fmt.Println(k, v)
	}
	fmt.Println("------")

	for k := range m{
		fmt.Println(k)
	}
	fmt.Println("------")

	for v := range m{
		fmt.Println(v)
	}
}
```