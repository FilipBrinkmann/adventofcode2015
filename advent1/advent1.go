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
	for _,v := range(bytes) {
		switch command := string(v); command {
		case "(":
			floor += 1
		case ")":
			floor -= 1
		default:
			panic("encountered unknown command!")
		}
	}
	fmt.Println(floor)
}

