package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var leftList []int
	var rightList []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		if len(parts) != 2 {
			log.Printf("Skipping invalid line: %s", scanner.Text())
			continue
		}

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("Error converting string %s to int: %v", parts[0], err)
			continue
		}

		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Printf("Error converting string %s to int: %v", parts[1], err)
			continue
		}

		leftList = append(leftList, num1)
		rightList = append(rightList, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDistance := 0
	r_m := make(map[int]int)
	for i := range leftList {
		distance := rightList[i] - leftList[i]
		if distance < 0 {
			distance *= -1
		}
		totalDistance += distance
		r_m[rightList[i]] += 1
	}

	// Part 1
	fmt.Printf("Total Distance: %d\n", totalDistance) // 1530215
	similarityScore := 0
	for i := range leftList {
		similarityScore += (leftList[i] * r_m[leftList[i]])
	}
	// Part 2
	fmt.Printf("Similarity Score: %d\n", similarityScore) // 26800609
}
