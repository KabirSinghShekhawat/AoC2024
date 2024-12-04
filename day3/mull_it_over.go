package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var buffer []byte
	for scanner.Scan() {
		buffer = append(buffer, scanner.Bytes()...)
	}
	input := string(buffer)
	prefix := "mul("
	l, r := 0, 0
	tokens := make([]string, 0)
	do := "do()"
	dont := "don't()"

	for i, ch := range input {
		if ch != rune('m') && ch != rune('d') {
			continue
		}
		l = i
		r = i
	tokenLoop:
		for r < len(input) {
			if ch == rune('d') {
				r2, err := peekToken(&input, l, dont)
				if err == nil {
					tokens = append(tokens, string(input[l:r2+1]))
					break tokenLoop
				}

				r2, err = peekToken(&input, l, do)
				if err == nil {
					tokens = append(tokens, string(input[l:r2+1]))
					break tokenLoop
				}
			}
			if prefix[(r-l)%len(prefix)] == input[r] {
				r++
			} else if (r-1)-l+1 == len(prefix) {
				if !unicode.IsDigit(rune(input[r])) {
					break tokenLoop
				}
				firstHalfParsed := false
				for rune(input[r]) != rune(')') {
					if unicode.IsDigit(rune(input[r])) {
						r++
					} else if !firstHalfParsed && rune(input[r]) == rune(',') {
						firstHalfParsed = true
						r++
					} else {
						break tokenLoop // this is not a valid token
					}
				}
				if r < len(input) && firstHalfParsed && rune(input[r]) == rune(')') {
					tokens = append(tokens, string(input[l:r+1]))
				}
				break tokenLoop
			} else {
				break tokenLoop
			}
		}
	}

	f2, err := os.Create("parsed_tokens.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	ans := 0
	ansPartTwo := 0
	mulEnable := true
	for i, t := range tokens {
		f2.WriteString(fmt.Sprintf("Token[%d]: %s\n", i, t))
		tokenLength := len(t)
		switch tokenLength {
		case len(do):
			mulEnable = true
		case len(dont):
			mulEnable = false
		default:
			res, err := parseMul(&t, &i, true)
			if err == nil {
				ans += res
			}
			res, err = parseMul(&t, &i, mulEnable)
			if err == nil {
				ansPartTwo += res
			}
		}
	}

	fmt.Println("Total valid instructions: ", len(tokens))
	fmt.Printf("Total sum = %d\n", ans)                   // 179571322
	fmt.Printf("Total sum (part two) = %d\n", ansPartTwo) // 103811193
}

func peekToken(input *string, l int, tok string) (int, error) {
	r := l
	for r < len(*input) && r-l+1 < len(tok) {
		if (*input)[r] != tok[(r-l)] {
			return l, errors.New("Invalid token")
		}
		r++
	}
	if r-l+1 == len(tok) && string((*input)[l:r+1]) == tok {
		return r, nil
	}
	return l, errors.New("token not found")
}

func parseMul(token *string, tokenNumber *int, mulEnable bool) (int, error) {
	if !mulEnable {
		return 0, nil
	}
	prefixLength := len("mul(")
	closingBracketIndex := (len(*token) - prefixLength)
	numPair := (*token)[prefixLength : prefixLength+closingBracketIndex-1]
	nums := strings.Split(numPair, ",")
	if len(nums) != 2 {
		fmt.Printf("Token[%d]: Expected 2 numbers, got %d\n", *tokenNumber, len(nums))
		fmt.Printf("Val: %v\n", nums)
		return 0, errors.New("Invalid token")
	}
	n1, err := strconv.Atoi(nums[0])
	if err != nil {
		fmt.Printf("[%d]: Failed to parse string %s as int\n", *tokenNumber, nums[0])
	}

	n2, err := strconv.Atoi(nums[1])
	if err != nil {
		fmt.Printf("[%d]: Failed to parse string %s as int\n", *tokenNumber, nums[1])
	}

	return n1 * n2, nil
}
