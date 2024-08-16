package main

import "fmt"

// 定数
const Pi = 3.14

const (
	URL = "http://xxx.com"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota		// 整数の連番を返す
	c1
	c2
)

func main(){
	fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(c0, c1, c2)
}