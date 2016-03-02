package main

import (
	"fmt"
	"strconv"
)

// --- Day 10: Elves Look, Elves Say ---

// Today, the Elves are playing a game called look-and-say.
// They take turns making sequences by reading aloud the previous sequence and using that reading as
// the next sequence. For example, 211 is read as "one two, two ones",
// which becomes 1221 (1 2, 2 1s).

// Look-and-say sequences are generated iteratively,
// using the previous value as input for the next step.
// For each step, take the previous value, and replace each run of digits (like 111) with the
// number of digits (3) followed by the digit itself (1).

// For example:

// 1 becomes 11 (1 copy of digit 1).
// 11 becomes 21 (2 copies of digit 1).
// 21 becomes 1211 (one 2 followed by one 1).
// 1211 becomes 111221 (one 1, one 2, and two 1s).
// 111221 becomes 312211 (three 1s, two 2s, and one 1).
// Starting with the digits in your puzzle input, apply this process 40 times.
// What is the length of the result?

// Your puzzle input is 1113222113.

var input string = "1113222113"

func main() {
	// Part one has 40 iterations, part two has 50
	//iterations := 40
	iterations := 50
	part_one_iterations := 40
	next_sequence := input
	for i := 0; i < iterations; i++ {
		// fmt.Println(next_sequence)
		fmt.Println("iteration", i)
		current_sequence := next_sequence
		next_sequence = ""
		// iterate over all sequences
		current_pointer := 0
		for current_pointer < len(current_sequence) {
			// time.Sleep(100 * time.Millisecond)
			// transform each sequence into string
			end_of_this_sequence := endOfCurrentSequence(current_sequence, current_pointer)
			// fmt.Println(end_of_this_sequence)
			new_subsequence := transformSequence(current_sequence[current_pointer:end_of_this_sequence])
			// fmt.Println(new_subsequence)
			// concatenate strings and add to next_sequence
			next_sequence += new_subsequence
			current_pointer = end_of_this_sequence
		}
		if i == part_one_iterations || i == iterations-1 {
			fmt.Println("Sequence has length ", len(next_sequence))
		}
	}
}

// Returns the index of the field right of the end of the current sequence. If, e.g., the current sequence
// has length 1, it returns the current pointer+1.
func endOfCurrentSequence(current_sequence string, current_pointer int) int {
	if current_pointer == len(current_sequence)-1 {
		return current_pointer + 1
	}
	sequence_char := current_sequence[current_pointer]
	for i := current_pointer + 1; i < len(current_sequence); i++ {
		if sequence_char != current_sequence[i] {
			return i
		}
	}
	return len(current_sequence) // same sequence char until end of string
}

// a sequence consists of a number of the same number literals
func transformSequence(subsequence string) string {
	character := string(subsequence[0])
	amount := len(subsequence)
	return strconv.Itoa(amount) + character
}
