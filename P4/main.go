package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the N")
	Num, _ := reader.ReadString('\n')
	Num = strings.Replace(Num, "\n", "", -1)
	IntNum, _ := strconv.Atoi(Num)
	K := IntNum / 1000
	G := IntNum % 1000
	fmt.Printf("the number is %v Kilogram and %v gram\n", K, G)
}
