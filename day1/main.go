package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func requirements(weight int) int {
	return int(weight/3) - 2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	flag.Parse()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		require := requirements(weight)

		for ok := true; ok; ok = require > 0 {
			sum = sum + require
			require = requirements(require)
		}

	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
