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
	fmt.Println("Please enter the number:")
	Num, _ := reader.ReadString('\n')
	Num = strings.Replace(Num, "\n", "", -1)
	IntNum, _ := strconv.Atoi(Num)
	A := IntNum / 10
	B := IntNum % 10
	Sum := A + B
	fmt.Printf("the sum is %v \n", Sum)
}
