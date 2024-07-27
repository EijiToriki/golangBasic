package main

import (
	"fmt"
	"strconv"
)

// 型

func main(){
	// 整数型
	var i int = 100
	var i2 int64 = 120

	fmt.Println(i + 50)
	// fmt.Println(i + i2)	// エラー

	// 型の調査
	fmt.Printf("$T\n", i2)

	// 型変換
	fmt.Printf("%T\n", int32(i2))
	fmt.Println(i + int(i2))


	// 浮動小数点数型
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)	// 普通は使わない
	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)


	// 論理値型
	var t, f bool = true, false
	fmt.Println(t, f)


	// 文字列型
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	test
		test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))


	// byte型
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))


	// 配列型（固定長）
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0], arr2[1], arr2[2])
	arr2[2] = "j"
	fmt.Println(arr2[0], arr2[1], arr2[2])

	// var arr5 [4]int
	// arr5 = arr1		// エラー

	fmt.Println(len(arr1))


	// interface型
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	// x = 2
	// fmt.Println(x + 3)		// エラー：演算の対象として利用できない


	// 型変換
	var i_tc int = 1
	fl64_tc := float64(i_tc)
	fmt.Println(fl64_tc)
	fmt.Printf("i = %T\n", i_tc)
	fmt.Printf("fl64 = %T\n", fl64_tc)

	i2_tc := int(fl64_tc)
	fmt.Printf("i2 = %T\n", i2_tc)

	fl_tc := 10.5
	i3_tc := int(fl_tc)
	fmt.Printf("i3 = %T\n", i3_tc)
	fmt.Println(i3_tc)

	var s_tc string = "100"
	fmt.Printf("s = %T\n", s_tc)
	fmt.Println(s_tc)

	i_s_tc, err := strconv.Atoi(s_tc)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(i_s_tc)
		fmt.Printf("i_s_tc = %T\n", i_s_tc)
	}

	var i_s_tc2 int = 200
	s2_tc := strconv.Itoa(i_s_tc2)
	fmt.Printf("s2 = %T\n", s2_tc)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Print(b)

	h2 := string(b)
	fmt.Println(h2)

}