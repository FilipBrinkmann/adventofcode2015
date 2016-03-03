package main

import (
	"fmt"
	"strings"
)

// --- Day 11: Corporate Policy ---

// Santa's previous password expired, and he needs help choosing a new one.

// To help him remember his new password after the old one expires, Santa has devised a method of
// coming up with a password based on the previous one. Corporate policy dictates that passwords
// must be exactly eight lowercase letters (for security reasons), so he finds his new password by
// incrementing his old password string repeatedly until it is valid.

// Incrementing is just like counting with numbers: xx, xy, xz, ya, yb, and so on. Increase the
// rightmost letter one step; if it was z, it wraps around to a, and repeat with the next letter to
// the left until one doesn't wrap around.

// Unfortunately for Santa, a new Security-Elf recently started, and he has imposed some additional
// password requirements:

// Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and
// so on, up to xyz. They cannot skip letters; abd doesn't count.

// Passwords may not contain the
// letters i, o, or l, as these letters can be mistaken for other characters and are therefore
// confusing.

// Passwords must contain at least two different, non-overlapping pairs of letters, like
// aa, bb, or zz. For example:

// hijklmmn meets the first requirement (because it contains the straight hij) but fails the second
// requirement requirement (because it contains i and l). abbceffg meets the third requirement
// (because it repeats bb and ff) but fails the first requirement. abbcegjk fails the third
// requirement, because it only has one double letter (bb). The next password after abcdefgh is
// abcdffaa. The next password after ghijklmn is ghjaabcc, because you eventually skip all the
// passwords that start with ghi..., since i is not allowed. Given Santa's current password (your
// puzzle input), what should his next password be?

// Your puzzle input is hepxcrrq.
func main() {
	password := "hepxcrrq"
	for !isAllowedPassword(&password) {
		password = *increment(&password)
	}
	fmt.Println(password)
}

func isAllowedPassword(password *string) bool {

	// first test
	passed_first_test := false
	if len(*password) < 3 {
		return false
	}
	for i := 0; i < len(*password)-3; i++ {
		if (*password)[i]+1 == (*password)[i+1] && (*password)[i+1]+1 == (*password)[i+2] {
			passed_first_test = true
			break
		}
	}
	if !passed_first_test {
		return false
	}
	// second test
	if strings.ContainsAny(*password, "i o l") {
		return false
	}
	return true
}

func increment(password *string) *string {
	bytes := []byte(*password)
	carry := false
	overflow := false
	for i := len(*password) - 1; i >= 0; i-- {
		if carry == false && i != len(*password)-1 {
			// nothing to do anymore
			break
		}
		if bytes[i] == 'z' {
			if i == 0 {
				// corner case: new string is one character longer now
				// in this case, we will simply add prepend an "a" before returning
				overflow = true
				bytes[i] = 'a'
			} else {
				carry = true
				bytes[i] = 'a'
			}
		} else {
			carry = false
			bytes[i] = bytes[i] + 1
		}
	}
	value := ""
	if overflow {
		value = string("a" + string(bytes))
	} else {
		value = string(bytes)
	}
	return &value
}
