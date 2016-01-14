package main

import  (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("input1.txt")
	if err != nil {
		panic(err)
	}

	floor := 0
	first_occurrence := true
	for i,v := range(bytes) {
		switch command := string(v); command {
		case "(":
			floor += 1
		case ")":
			floor -= 1
		default:
			panic("encountered unknown command!")
		}
		if floor == -1 && first_occurrence  {
				defer fmt.Printf("Basement was entered upon the %dth command\n", i+1)
				first_occurrence = false
		}
	}
	defer fmt.Printf("Santa ended up on the %dth floor\n", floor)
}

