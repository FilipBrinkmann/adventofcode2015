package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"regexp"
	"strconv"
)

// --- Day 12: JSAbacusFramework.io ---

// Santa's Accounting-Elves need help balancing the books after a recent order. Unfortunately, their
// accounting software uses a peculiar storage format. That's where you come in.

// They have a JSON document which contains a variety of things: arrays ([1,2,3]), objects ({"a":1,
// "b":2}), numbers, and strings. Your first job is to simply find all of the numbers throughout the
// document and add them together.

// For example:

// [1,2,3] and {"a":2,"b":4} both have a sum of 6. [[[3]]] and {"a":{"b":4},"c":-1} both have a
// sum of 3.  {"a":[-1,1]} and [-1,{"a":1}] both have a sum of 0. [] and {} both have a sum of 0.
// You will not encounter any strings containing numbers.

// What is the sum of all numbers in the document?

func main() {
	// part one
	sum := 0
	bytes, err := ioutil.ReadFile("input12.txt")
	if err != nil {
		panic("could not read file")
	}
	re, err := regexp.Compile("-?[0-9]+")
	if err != nil {
		panic("error matching regex")
	}
	matches := re.FindAll(bytes, -1)
	for _, match := range matches {
		current_number, err := strconv.Atoi(string(match))
		if err != nil {
			panic("Error parsing number " + strconv.Itoa(current_number))
		} else {
			sum += current_number
		}
	}
	fmt.Println("sum: ", strconv.Itoa(sum))

	// PART TWO
	// Uh oh - the Accounting-Elves have realized that they double-counted everything
	// red.

	// Ignore any object (and all of its children) which has any property with the value "red".
	// Do this only for objects ({...}), not arrays ([...]).

	// [1,2,3] still has a sum of 6. [1,{"c":"red","b":2},3] now has a sum of 4, because the middle object
	// is ignored. {"d":"red","e":[1,2,3,4],"f":5} now has a sum of 0, because the entire structure is
	// ignored. [1,"red",5] has a sum of 6, because "red" in an array has no effect.
	json.
}
