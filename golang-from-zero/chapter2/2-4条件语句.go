package main

import (
	"fmt"
	"os"
)

func openFileCondition1() {
	content, err := os.ReadFile("./abc.txt")
	if err != nil {
		fmt.Println("connot print file content:", err)
	} else {
		fmt.Printf("%q\n", content)
	}
}

func openFileCondition2() {
	if content, err := os.ReadFile("./abc.txt"); err != nil {
		fmt.Println("connot print file content:", err)
	} else {
		fmt.Println(string(content))
	}
}

func grade(score int) string {
	var g string
	switch {
	case score < 0 || score > 100:
		panic("error score")
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	openFileCondition1()
	openFileCondition2()
	fmt.Println(
		grade(0),
		grade(59),
		grade(70),
		grade(85),
		// grade(101),
	)
}
