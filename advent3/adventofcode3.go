package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	bytes, err := ioutil.ReadFile("input3.txt")
	if err != nil {
		panic(err)
	}

	x := 0
	y := 0
	houses_visited := 1 // 0,0 is already visited
	visited_coordinates := make(map[string]bool)
	visited_coordinates["0,0"] = true
	for _, byte := range(bytes) {
		valid_input := true
		switch direction := string(byte); direction {
			case "^":
				y += 1
			case "<":
				x -= 1
			case ">":
				x += 1
			case "v":
				y -= 1
			default:
				valid_input = false
		}
		if valid_input {
			coord := strings.Join([]string{strconv.Itoa(x),strconv.Itoa(y)}, ",")
			_, ok := visited_coordinates[coord]
			if ! ok {
				visited_coordinates[coord] = true
				houses_visited++
			} else {
				visited_coordinates[coord] = false
			}
		}
	}
	fmt.Printf("Santa visited %d houses", houses_visited)
}