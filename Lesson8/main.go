package main

import (
	"fmt"
	"os"
	"time"
	// "strconv"
)

func init(){
	fmt.Println("init")
}

// if
func main(){
	//条件分岐
	// a := 0
	// if a == 2 {
	// 	fmt.Println("two")
	// }else if a == 1{
	// 	fmt.Println("one")
	// }else{
	// 	fmt.Println("I don`t know")
	// }

	// if b:=100; b == 100{
	// 	fmt.Println("one hundred")
	// }

	// x := 0
	// if x:= 2; true{
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// エラーハンドリング
	// var s string = "A"

	// i, err := strconv.Atoi(s)
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// fmt.Printf("i = %T\n", i)

	// for文
	// i2 := 0
	// for{
	// 	i2 ++
	// 	if i2 == 3{
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// point := 0
	// for point < 10{
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i:=0; i<10; i++{
	// 	if i==3{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1, 2, 3}
	// for i:=0; i<len(arr); i++{
	// 	fmt.Println(arr[i])
	// }
	// for i, v := range arr{
	// 	fmt.Println(i, v)
	// }

	// sl := []string{"Python", "PHP", "GO"}
	// for i, v := range sl{
	// 	fmt.Println(i, v)
	// }

	// m := map[string]int{"apple": 100, "banana": 200}
	// for k, v := range m{
	// 	fmt.Println(k, v)
	// }

	// switch
	// n := 5
	// switch n{
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don`t know")
	// }

	// switch n := 2; n{
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don`t know")
	// }

	// n := 6
	// switch {
	// case n > 0 && n < 4:
	// 	fmt.Println("0 < n < 4")
	// case n > 3 && n < 7:
	// 	fmt.Println("3 < n < 7")
	// // case 4
	// // 	fmt.Println(4)	エラー
	// }

	// 型スイッチ
	// var x interface{} = 3
	// i, isInt := x.(int)
	// fmt.Println(i, isInt)

	// f, isFloat64 := x.(float64)
	// fmt.Println(f, isFloat64)

	// if x == nil{
	// 	fmt.Println("None")
	// }else if i, isInt := x.(int); isInt {
	// 	fmt.Println(i, "x is Int")
	// }else if s, isString := x.(string); isString {
	// 	fmt.Println(s, isString)
	// }else{
	// 	fmt.Println("I don`t know")
	// }

	// switch x.(type) {
	// case int:
	// 	fmt.Println("int")
	// case string:
	// 	fmt.Println("string")
	// default:
	// 	fmt.Println("I don`t know")
	// }

	// switch v := x.(type) {
	// case int:
	// 	fmt.Println(v, "int")
	// case string:
	// 	fmt.Println(v, "string")
	// default:
	// 	fmt.Println(v, "I don`t know")
	// }

	anything(1)
	anything("aaa")

	// ラベル付きfor
	// Loop:
	// 	for{
	// 		for{
	// 			for{
	// 				fmt.Println("Start")
	// 				break Loop
	// 			}
	// 			fmt.Println("処理しない")
	// 		}
	// 		fmt.Println("処理しない")
	// 	}
	// fmt.Println("END")

	// Loop:
	// 	for i:=0; i<3; i++{
	// 		for j:=1; j<3; j++{
	// 			if j>1 {
	// 				continue Loop
	// 			}
	// 			fmt.Println(i, j, i*j)
	// 		}
	// 		fmt.Println("処理しない")
	// 	}

	// defer
	TestDefer()

	// main関数の終了時に走る
	// defer func(){
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	// RunDefer()

	// リソースの開放処理
	file, err := os.Create("test.txt")
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))

	// panic recover
	// panic("runtime error")
	// fmt.Println("Start")

	// defer func ()  {
	// 	if x := recover(); x != nil{
	// 		fmt.Println(x)
	// 	}
	// }()
	// panic("runtime error")
	// fmt.Println("start")

	// gorutin
	// go sub()

	// for{
	// 	fmt.Println("Main loop")
	// 	time.Sleep(200 * time.Millisecond)
	// }

}

func anything(a interface{}){
	switch v := a.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println(v, "I don`t know")
	}
}


func TestDefer(){
	defer fmt.Println("END")
	fmt.Println("Start")
}

func RunDefer(){
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func sub(){
	for {
		fmt.Println("Sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}