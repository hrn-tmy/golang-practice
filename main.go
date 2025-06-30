package main

import "fmt"

func main () {
	var a any
	fmt.Println(a) // 初期値：nil

	var b bool
	fmt.Println(b) // 初期値：false

	var c byte
	fmt.Println(c) // 初期値：0

	var d complex64
	fmt.Println(d) // 初期値：(0+0i)

	var e float32
	fmt.Println(e) // 初期値：0(float64も同じ)

	var f int
	fmt.Println(f) // 初期値：0(int8,int16,int32,int64も同じ)

	var g rune
	fmt.Println(g) // 初期値：0

	var h string
	fmt.Println(h) // 初期値：(空文字)

	var i uint
	fmt.Println(i) // 初期値：0(uint8,uint16,uint32,uint64も同じ)
	
	var j uintptr
	fmt.Println(j) // 初期値：0
}