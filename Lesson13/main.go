package main

import (
	"fmt"
	f "fmt"
	. "fmt"
	"foo/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

// エラー
// func Do(s string) (b string){
// 	var b string = s
// 	return b
// }

func main() {
	fmt.Println(foo.Max)
	f.Println(foo.ReturnMin())
	Println()

	fmt.Println(appName())
	// fmt.Println(AppName, Version)	//エラー
}