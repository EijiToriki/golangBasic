package main

import (
	"fmt"
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

}


func Sum(s ...int) int{
	n := 0
	for _, v := range s{
		n += v
	}
	return n
}