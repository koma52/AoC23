package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"day-01/part1"
	"day-01/part2"
)

func main() {
	f, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var words = []string{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	fmt.Println(part1.Part1(words))
	fmt.Println(part2.Part2(words))
}
