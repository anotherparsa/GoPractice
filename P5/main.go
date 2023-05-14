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
	fmt.Println("Please enter the N:")
	Num, _ := reader.ReadString('\n')
	Num = strings.Replace(Num, "\n", "", -1)
	IntNum, _ := strconv.Atoi(Num)
	Hours := IntNum / 3600
	PlaceHolder := IntNum % 3600
	Minutes := PlaceHolder / 60
	Seconds := PlaceHolder % 60
	fmt.Printf("the %v seconds is %v hours and %v minutes and %v seconds\n", IntNum, Hours, Minutes, Seconds)

}
