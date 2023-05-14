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
	fmt.Println("Please enter the first side:")
	FirstSide, _ := reader.ReadString('\n')
	FirstSide = strings.Replace(FirstSide, "\n", "", -1)
	IntFirstSide, _ := strconv.Atoi(FirstSide)
	fmt.Println("Please enter the second side:")
	SecondSide, _ := reader.ReadString('\n')
	SecondSide = strings.Replace(SecondSide, "\n", "", -1)
	IntSecondSide, _ := strconv.Atoi(SecondSide)
	fmt.Println("Please enter the third side:")
	ThirdSide, _ := reader.ReadString('\n')
	ThirdSide = strings.Replace(ThirdSide, "\n", "", -1)
	IntThirdSide, _ := strconv.Atoi(ThirdSide)
	if (IntFirstSide * IntFirstSide) == (IntSecondSide*IntSecondSide)+(IntThirdSide*IntThirdSide) {
		fmt.Println("this is a right triangle")
	} else if (IntSecondSide * IntSecondSide) == (IntFirstSide*IntFirstSide)+(IntThirdSide*IntThirdSide) {
		fmt.Println("this is a right triangle")
	} else if (IntThirdSide * IntThirdSide) == (IntFirstSide*IntFirstSide)+(IntSecondSide*IntSecondSide) {
		fmt.Println("this is a right triangle")
	} else {
		fmt.Println("this is not a right triangle")
	}
}
