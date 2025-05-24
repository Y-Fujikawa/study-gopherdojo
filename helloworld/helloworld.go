package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	// 変数定義いろいろ
	var n int = 100
	var s string = "hello"
	var b bool = true
	var f float64 = 3.14
	var arr []int = []int{1, 2, 3}
	var m map[string]int = map[string]int{"apple": 1, "banana": 2}

	fmt.Println(n, s, b, f, arr, m)

	// 変数定義省略
	// 省略形は関数内でのみ使用可能
	n2 := 100
	s2 := "hello"
	b2 := true
	f2 := 3.14
	arr2 := []int{1, 2, 3}
	m2 := map[string]int{"apple": 1, "banana": 2}

	fmt.Println(n2, s2, b2, f2, arr2, m2)

	// 変数定義をまとめる
	var (
		n3 int     = 100
		s3 string  = "hello"
		b3 bool    = true
		f3 float64 = 3.14
	)

	fmt.Println(n3, s3, b3, f3)

	// 定数
	const c4 int = 100
	const s4 string = "hello"
	const b4 bool = true
	const f4 float64 = 3.14

	fmt.Println(c4, s4, b4, f4)

	// 型変換
	var i5 int = 100
	var f5 float64 = 3.14
	var s5 string = "hello"

	fmt.Println(i5, f5, s5)

	// 型変換
	var i6 int = 100
	var f6 float64 = 3.14
	var s6 string = "hello"

	fmt.Println(i6, f6, s6)

	// 変数のゼロ値
	// 明示的な初期化をしなくてもゼロ値が設定される
	var n7 int
	var s7 string
	var b7 bool
	var f7 float64
	var arr7 []int
	var m7 map[string]int

	fmt.Println(n7, s7, b7, f7, arr7, m7)

	// 簡単な例
	// sum := 0
	// 以下のように書いても良い
	var sum int
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
