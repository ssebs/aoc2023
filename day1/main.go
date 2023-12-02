package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetDigits(str string) (digits int) {
	first, last := -1, -1

	// L=>R
	for _, c := range str {
		if num, err := strconv.Atoi(string(c)); err == nil {
			first = num
			break
		}
	}

	// R=>L
	for i := len(str) - 1; i >= 0; i-- {
		if num, err := strconv.Atoi(string(str[i])); err == nil {
			last = num
			break
		}
	}
	fmt.Printf("combined: %d%d", first, last)
	digits, _ = strconv.Atoi(fmt.Sprintf("%d%d", first, last))
	return digits
}

func main() {
	sum := 0

	f, err := os.Open("./day1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		digits := GetDigits(scanner.Text())
		fmt.Printf("%s => %d\n", scanner.Text(), digits)
		sum += digits
	}
	fmt.Println("Sum:", sum)
}
