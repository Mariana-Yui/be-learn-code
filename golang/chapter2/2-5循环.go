package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertBin(num int) string {
	result := ""
	for i := num; i > 0; i /= 2 {
		lsb := i % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	convertBin(3)
	convertBin(13)
	convertBin(826382)
	printFile("abc.txt")
	// forever()

}
