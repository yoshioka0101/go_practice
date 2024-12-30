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