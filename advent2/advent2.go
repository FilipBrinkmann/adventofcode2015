package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var total_area int = 0
	var total_ribbon int = 0
	var packets [][]int = make([][]int, 0)
	for _, line := range strings.Fields(string(bytes)) {
		lengths := strings.Split(line, "x")
		w, ok := strconv.Atoi(lengths[0])
		if ok != nil {
			panic("error converting")
		}
		l, ok := strconv.Atoi(lengths[1])
		if ok != nil {
			panic("error converting")
		}
		h, ok := strconv.Atoi(lengths[2])
		if ok != nil {
			panic("error converting")
		}
		dimensions := []int{w, l, h}
		sort.Ints(dimensions)
		packets = append(packets, dimensions)
	}
	for _, packet := range packets {
		fmt.Println(packet)
		areas := []int{packet[0] * packet[1], packet[1] * packet[2], packet[0] * packet[2]}
		sort.Ints(areas)
		area := 2*areas[0] + 2*areas[1] + 2*areas[2] + 	// packet surface
			areas[0] 									// extra paper of size of smallest side: areas[0] is the smallest
		ribbon_length := 2*packet[0] + 2*packet[1] + 	// ribbon wraps around two smallest faces
			packet[0]*packet[1]*packet[2] 				// bow: volume of package
		total_area += area
		total_ribbon += ribbon_length
	}
	fmt.Println("Total area of paper needed:", total_area)
	// correct solution of part 1: 1588178
	fmt.Println("Total ribbon length:", total_ribbon)
	// correct solution of part 2: 3783758
}
