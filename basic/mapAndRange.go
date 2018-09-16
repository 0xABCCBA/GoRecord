package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println(m)

	vk1 := m["k1"]
	fmt.Println("vk1: ", vk1)

	for k, v := range m {
		fmt.Println(k ,v)
	}

	fmt.Println("len: ", len(m))	//返回键值对个数

	delete(m, "k2")
	fmt.Println("map: ", m)

	v, exist := m["k1"]
	fmt.Println("k1v: ", v, exist)

	n := map[string]int{"key1": 12, "key2": 13}
	fmt.Println(n)

	str := "中文"
	for i, c := range str {
		fmt.Printf("%d, %q\n", i, c)
		fmt.Println(string(c))
	}
	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i])
	}

	str = "Hello,世界"
    for i := 0; i < len(str); i++ {
        ch := str[i]
        ctype:=reflect.TypeOf(ch)
        fmt.Printf("%s ",ctype)
    }
	fmt.Println()
	//unicode 遍历
    for _, ch1 := range str {
        ctype:=reflect.TypeOf(ch1)
        fmt.Printf("%s ",ctype)
    }
	fmt.Println()


}
