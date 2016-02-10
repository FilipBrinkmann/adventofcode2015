package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input3.txt")
	if err != nil {
		panic(err)
	}

	homes_visited_by_lonely_santa := 1

	// coordinates of lonely santa and the santa team
	x_santa_alone, x_santa, x_robo := 0, 0, 0
	y_santa_alone, y_santa, y_robo := 0, 0, 0

	visited_coordinates := make(map[string]int)
	visited_coordinates[strings.Join([]string{strconv.Itoa(x_santa), strconv.Itoa(y_santa)}, ",")] = 0
	visited_coordinates_alone := make(map[string]int)
	visited_coordinates_alone[strings.Join([]string{strconv.Itoa(x_santa_alone), strconv.Itoa(y_santa_alone)}, ",")] = 0

	// calculate coordinates
	for count, byte := range bytes {
		valid_input := true
		switch direction := string(byte); direction {
		case "^":
			if count%2 == 0 {
				y_santa += 1
			} else {
				y_robo += 1
			}
			y_santa_alone += 1
		case "<":
			if count%2 == 0 {
				x_santa -= 1
			} else {
				x_robo -= 1
			}
			x_santa_alone -= 1
		case ">":
			if count%2 == 0 {
				x_santa += 1
			} else {
				x_robo += 1
			}
			x_santa_alone += 1
		case "v":
			if count%2 == 0 {
				y_santa -= 1
			} else {
				y_robo -= 1
			}
			y_santa_alone -= 1
		default:
			valid_input = false
		}

		// update visited homes
		if valid_input {
			var x int
			var y int
			if count%2 == 0 {
				x = x_santa
				y = y_santa
			} else {
				x = x_robo
				y = y_robo
			}
			coordinates_alone := strings.Join([]string{strconv.Itoa(x_santa_alone), strconv.Itoa(y_santa_alone)}, ",")
			// fmt.Println(coordinates_alone)
			coordinates_team := strings.Join([]string{strconv.Itoa(x), strconv.Itoa(y)}, ",")
			_, ok := visited_coordinates_alone[coordinates_alone]
			_, ok_team := visited_coordinates[coordinates_team]
			if !ok { // no entry yet: house not visited
				homes_visited_by_lonely_santa++ // first part here
				visited_coordinates_alone[coordinates_alone] = 0
			}
			if !ok_team { // no entry yet: house not visited
				// second part
				if count%2 == 0 {
					visited_coordinates[coordinates_team] = 0 // santa
				} else {
					// Robosanta visited
					visited_coordinates[coordinates_team] = 1 // robosanta
				}
			}
		} else {
			panic("Erroneous input.")
		}
	}

	fmt.Printf("Santa visited %d homes\n", homes_visited_by_lonely_santa)
	fmt.Printf("Together with Robo-Santa, he visited %d homes.", len(visited_coordinates))
}

// Correct solutions are:
// 2565 (first part)
// 2639 (second part)
