package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// This is a slice containing all rules
var rules []func(string) bool

func main() {
	set_rules_part_one()
	nice_strings := 0
	bytes, err := ioutil.ReadFile("input5.txt")
	if err != nil {
		panic("error reading file")
	}
	// Part ONE
	fmt.Println("Part one:")
	for _, line := range strings.Fields(string(bytes)) {
		if isNice(line) {
			nice_strings++
		}
	}
	fmt.Printf("There are %d nice strings\n", nice_strings)
	fmt.Println("Hint: the correct answer is 255")
	// Part TWO
	fmt.Println("\nPart Two:")
	nice_strings = 0
	set_rules_part_two()
	for _, line := range strings.Fields(string(bytes)) {
		if isNice(line) {
			nice_strings++
		}
	}
	fmt.Printf("There are %d nice strings\n", nice_strings)
	fmt.Println("Hint: the correct answer is 55")
}

func set_rules_part_one() {
	rules = nil
	// Rule one: does not contain bad substrings
	rule := func(line string) bool {
		for _, sub := range []string{"ab", "cd", "pq", "xy"} {
			if strings.Contains(line, sub) {
				return false
			}
		}
		return true
	}
	rules = append(rules, rule)
	// Rule two: contains at least 3 vowels
	rule = func(line string) bool {
		vowels := 0
		characters := []byte(line)
		for _, v := range characters {
			if strings.Contains("aeiou", string(v)) {
				vowels++
			}
			if vowels == 3 {
				return true
			}
		}
		return false
	}
	rules = append(rules, rule)
	// Rule three: it contains at least one double letter
	rule = func(line string) bool {
		characters := []byte(line)
		for i := 0; i < (len(characters) - 1); i++ {
			if characters[i] == characters[i+1] {
				return true
			}
		}
		return false
	}
	rules = append(rules, rule)
}

func set_rules_part_two() {
	rules = nil
	// Rule one: It contains at least one letter which repeats with exactly one letter between
	// them, like xyx, abcdefeghi (efe), or even aaa.
	rule := func(line string) bool {
		if len(line) < 3 {
			return false
		}
		for i := 0; i < (len(line) - 2); i++ {
			if line[i] == line[i+2] {
				return true
			}
		}
		return false
	}
	rules = append(rules, rule)
	// Rule two: It contains a pair of any two letters that appears at least twice in the string
	// without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa
	// (aa, but it overlaps).
	rule = func(line string) bool {
		if len(line) < 4 {
			return false
		}
		for i := 0; i < (len(line) - 3); i++ {
			pattern := line[i : i+2]
			if strings.Contains(line[i+2:len(line)], pattern) {
				return true
			}
		}
		return false
	}
	rules = append(rules, rule)
}

func isNice(line string) bool {
	for _, rule := range rules {
		if rule(line) == false {
			return false
		}
	}
	return true
}
