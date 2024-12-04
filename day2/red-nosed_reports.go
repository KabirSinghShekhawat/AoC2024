package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isSafeReport determines if a report is safe based on two conditions:
// 1. Levels are either all increasing or all decreasing
// 2. Adjacent levels differ by at least 1 and at most 3
func isSafeReport(levels []int) bool {
	// Check if levels are increasing
	increasing := true
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		if diff < 1 || diff > 3 {
			increasing = false
			break
		}
	}

	// Check if levels are decreasing
	decreasing := true
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i] - levels[i+1]
		if diff < 1 || diff > 3 {
			decreasing = false
			break
		}
	}

	return increasing || decreasing
}

// isSafeReportPartTwo determines if an unsafe report can become safe if a single level is removed.
func isSafeReportPartTwo(levels []int) bool {
	for i := range levels {
		withoutLevelI := make([]int, 0, len(levels)-1)
		withoutLevelI = append(withoutLevelI, levels[:i]...)
		withoutLevelI = append(withoutLevelI, levels[i+1:]...)
		if isSafeReport(withoutLevelI) {
			return true
		}
	}
	return false
}

// countSafeReports counts the number of safe reports in the given list
func countSafeReports(reportList [][]int) int {
	safeCount := 0
	for _, report := range reportList {
		if isSafeReport(report) {
			safeCount++
		}
	}
	return safeCount
}

// countSafeReportsPartTwo counts the number of safe reports in the given list with dampening
func countSafeReportsPartTwo(reportList [][]int) int {
	safeCount := 0
	for _, report := range reportList {
		if isSafeReport(report) {
			safeCount++
		} else if isSafeReportPartTwo(report) {
			safeCount++
		}
	}

	return safeCount
}

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read reports from the file
	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line and convert to integers
		lineParts := strings.Fields(scanner.Text())
		report := make([]int, len(lineParts))
		for i, part := range lineParts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting number:", err)
				return
			}
			report[i] = num
		}
		reports = append(reports, report)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Count and print safe reports
	safeReportsCount := countSafeReports(reports)
	fmt.Printf("Number of safe reports: %d\n", safeReportsCount) // 224
	// Count and print safe reports part two
	safeReportsCountPartTwo := countSafeReportsPartTwo(reports)
	fmt.Printf("Number of safe reports with dampening: %d\n", safeReportsCountPartTwo) // 293
}
