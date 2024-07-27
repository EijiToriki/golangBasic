package main

import "fmt"

// 変数

// global := "global"		：エラー
var global string = "global"

func outer(){
	var s4 string = "outer"
	fmt.Println(s4)
}


func main(){
	// 明示的な定義
	var i int = 100
	fmt.Println(i)

	i = 150
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int			// 初期値 0
	var s3 string		// 初期値 ""
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 暗黙的な定義（関数内のみ）
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// i4 := 500		：エラー
	// i4 = "Hello" ：エラー


	// 関数呼び出し
	outer()
	// fmt.Println(s4)	：エラー　s4はouterの変数

	// var s5 string = "not use"		// 使わないとエラー

}