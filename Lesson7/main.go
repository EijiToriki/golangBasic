package main

import "fmt"

// 関数
func Plus(x int, y int) int {
	return x + y
}

func Div(x, y int)(int ,int){
	q := x / y
	r := x % y
	return q, r
}

func Double(price int)(result int){
	result = price * 2
	return
}

func Noreturn(){
	fmt.Println("No Return")
	return
}

// 関数を返す関数
func ReturnFunc() func(){
	return func(){
		fmt.Println(("I`m a function"))
	}
}

// 関数を引数に取る関数
func CallFunction(f func()){
	f()
}

// クロージャーの実装
func Later() func(string) string{
	var store string
	return func(next string) string{
		s := store
		store = next
		return s
	} 
}

// ジェネレータの実装
func integers() func() int{
	i := 0
	return func() int {
		i++
		return i
	} 
}

func main(){
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 4)
	fmt.Println(i2, i3)

	i4, _ := Div(9, 4)
	fmt.Println(i4)

	i5 := Double(1000)
	fmt.Println(i5)

	Noreturn()


	// 無名関数
	f := func(x, y int) int{
		return x + y
	}
	j := f(1, 2)
	fmt.Println(j)

	j2 := func(x, y int) int{
		return x + y
	}(1, 2)
	fmt.Println(j2)

	// 関数を返す関数
	g := ReturnFunc()
	g()

	// 関数を引数に取る関数
	CallFunction(func ()  {
		fmt.Println("I`m a function 2")
	})

	// クロージャ
	f_closer := Later()
	fmt.Println(f_closer("Hello"))
	fmt.Println(f_closer("My"))
	fmt.Println(f_closer("Name"))
	fmt.Println(f_closer("is"))
	fmt.Println(f_closer("Golang"))

	// ジェネレータ
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())


}