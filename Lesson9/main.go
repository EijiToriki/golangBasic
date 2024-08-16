package main

import (
	"fmt"
	"time"
)

func main() {
	// スライス
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)
	fmt.Println(sl5[0])
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:2])
	fmt.Println(sl5[2:])
	fmt.Println(sl5[:])
	fmt.Println(sl5[len(sl5)-1])
	fmt.Println(sl5[1 : len(sl5)-1])

	sl5 = append(sl5, 300)
	fmt.Println(sl5)

	sl5 = append(sl5, 400, 500, 600)
	fmt.Println(sl5)

	sl6 := make([]int, 5)
	fmt.Println(sl6)
	fmt.Println(len(sl6))
	fmt.Println(cap(sl6))

	sl7 := make([]int, 5, 10)
	fmt.Println(len(sl7))
	fmt.Println(cap(sl7))

	sl7 = append(sl7, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl7))
	fmt.Println(cap(sl7))

	sl8 := []int{100, 200, 300, 400, 500}
	sl9 := make([]int, 5, 10)
	n := copy(sl9, sl8)

	fmt.Println(n, sl9)

	sl10 := []string{"A", "B", "C"}
	for i, v := range sl10{
		fmt.Println(i, v)
	}

	for i:= 0; i < len(sl10); i++ {
		fmt.Println(sl10[i])
	}
	
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3))
	sl11 := []int{11, 22, 33}
	fmt.Println(Sum(sl11...))

	// マップ
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)
	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3])

	s, ok := m4[4]
	if !ok{
		fmt.Println("error")
	}
	fmt.Println(s, ok)

	m4[2] = "US"
	m4[3] = "China"
	fmt.Println(m4)

	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))

	m_for := map[string]int{
		"Apple": 100, 
		"Banana": 200,
	}

	for k, v := range m_for {
		fmt.Println(k, v)
	}

	// channel
	// 送受信どちらも
	var ch1 chan int
	// 受信専用
	// var ch2 <-chan int
	// // 送信専用
	// var ch3 chan<- int

	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1
	fmt.Println("len", len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println(len(ch3))

	i := <- ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))
	i2 := <- ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	// ch3 <- 6	// チャネルのサイズオーバでデッドロック

	ch4 := make(chan int)
	ch5 := make(chan int)
	// fmt.Println(<-ch4)	// エラー
	
	go receiver(ch4)
	go receiver(ch5)

	i3 := 0
	for i3 < 10{
		ch4 <- i3
		ch5 <- i3
		time.Sleep(50*time.Millisecond)
		i3++
	}

	ch1_close := make(chan int, 2)
	// close(ch1_close)
	// // ch1_close <- 1	//エラー
	// // fmt.Println(<-ch1_close)	// 受信はできる

	// i4, ok := <-ch1_close
	// fmt.Println(i4, ok)

	go receiver_close("1.gorutin", ch1_close)
	go receiver_close("2.gorutin", ch1_close)
	go receiver_close("3.gorutin", ch1_close)
	
	i5 := 0
	for i5 < 10{
		ch1_close <- i5
		i5++
	}
	close(ch1_close)
	time.Sleep(3*time.Second)


	ch1_for := make(chan int, 3)
	ch1_for <- 1
	ch1_for <- 2
	ch1_for <- 3
	close(ch1_for)		// これを入れないとエラー
	for j := range ch1_for{
		fmt.Println(j)
	}

	ch1_select := make(chan int, 2)
	ch2_select := make(chan string, 2)

	// ch1_select <- 1
	// ch2_select <- "A"

	select{
		case v1 := <-ch1_select:
			fmt.Println(v1 + 1000)
		case v2 := <-ch2_select:
			fmt.Println(v2 + "!?")
		default:
			fmt.Println("どちらでもない")
	}

	ch3_select := make(chan int)
	ch4_select := make(chan int)
	ch5_select := make(chan int)

	go func ()  {
		for {
			i_se := <- ch3_select
			ch4_select <- i_se * 2
		}
	}()

	go func ()  {
		for {
			i2_se := <- ch4_select
			ch5_select <- i2_se - 1
		}
	}()

	n2 := 0
	for {
		select {
			case ch3_select <- n2:
				n2++
			case i3_se := <-ch5_select:
				fmt.Println("received", i3_se)
		}
		if n2 > 100{
			break
		}
	}

}


func Sum(s ...int) int{
	n := 0
	for _, v := range s{
		n += v
	}
	return n
}

func receiver(c chan int){
	for{
		i := <-c
		fmt.Println(i)
	}
}

func receiver_close(name string, ch chan int){
	for{
		i, ok := <-ch
		if !ok{
			break
		}else{
			fmt.Println(i, name)
		}
	}
	fmt.Println(name + "END")
}