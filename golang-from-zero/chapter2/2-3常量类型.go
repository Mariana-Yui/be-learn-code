package main

import "fmt"

func consts() {
	const a, b = 3, 4
	const str = "Mariana"
	fmt.Println(a, b, str)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	const name = iota // iota只能用来赋值
	fmt.Println(cpp, java, python, golang, name)
}

func sizeEnums() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	consts()
	enums()
	sizeEnums()
}
