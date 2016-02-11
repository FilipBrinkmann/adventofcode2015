// Advent of Code 2015, Day 4
// Go solution
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

var key string = "ckczppom"
var part_one_prefix string = "00000"
var part_two_prefix string = "000000"
var MAX_ITERATIONS = 10000000

func serial_hashing() {
	var i int = 0
	one_found := false
	two_found := false
	for !(one_found && two_found) {
		hash := md5.New()
		serial := strconv.Itoa(i)
		// fmt.Println(key + serial)
		// fmt.Println(key)
		io.WriteString(hash, key+serial)
		thing := fmt.Sprintf("%x", hash.Sum(nil))
		if i >= MAX_ITERATIONS {
			panic("Aborting. No hash found")
		}
		if strings.HasPrefix(thing, part_one_prefix) && !one_found {
			// correct answer is 117946
			fmt.Println("Part ONE:")
			fmt.Println("Lucky number is " + strconv.Itoa(i))
			fmt.Println("Hash is " + thing)
			one_found = true
		}
		if strings.HasPrefix(thing, part_two_prefix) && !two_found {
			// correct answer is 3938038
			fmt.Println("Part TWO:")
			fmt.Println("Lucky number is " + strconv.Itoa(i))
			fmt.Println("Hash is " + thing)
			two_found = true
		}
		i++
	}
}

func search_hash(prefix string, start int, step int, c chan int) {
	for i := start; ; i = i + step {
		hash := md5.New()
		serial := strconv.Itoa(i)
		io.WriteString(hash, key+serial)
		thing := fmt.Sprintf("%x", hash.Sum(nil))
		if strings.HasPrefix(thing, prefix) {
			// correct answer is
			// 117946 for first part
			// and 3938038 for second
			c <- i
			return
		}
	}
}

func parallel_hashing() {
	var concurrency = 2
	one := make(chan int, 1)
	two := make(chan int, 1)
	timeout := time.After(10 * time.Second)
	one_found := false
	two_found := false
	for i := 0; i < concurrency; i++ {
		go search_hash(part_one_prefix, i, concurrency, one)
		go search_hash(part_two_prefix, i, concurrency, two)
	}
	for !(one_found && two_found) {
		select {
		case res_one := <-one:
			if res_one > 0 {
				one_found = true
				fmt.Println("Found ", res_one)
			}
		case res_two := <-two:
			if res_two > 0 {
				two_found = true
				fmt.Println("Found ", res_two)
			}
		case <-timeout:
			fmt.Println("timeout!")
			return
		}
	}
}

func main() {
	start := time.Now()
	serial_hashing()
	serial_duration := time.Since(start)

	start = time.Now()
	parallel_hashing()
	parallel_duration := time.Since(start)

	fmt.Printf("Serial execution took %.2f seconds.\n", serial_duration.Seconds())
	fmt.Printf("Parallel execution took %.2f seconds.\n", parallel_duration.Seconds())
}
