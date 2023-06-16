package main

import (
	"chatcser/pkg/plink/router2"
	"fmt"
)

func main() {
	tr := router2.NewPrefixTree()
	tr.Insert("/user/:nam", func(c any) {
		fmt.Println("123")
	})
	tr.Insert("/user/:aaa/:aa", func(c any) {
		fmt.Println("123")
	})

	tr.Insert("/user/:bb/:cc", func(c any) {
		fmt.Println("123")
	})
	n, m := tr.Search("/user/zhang/aa")
	n.Handler("33")
	println(n.Param)
	fmt.Printf("%v", m)

}
