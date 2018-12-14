package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	p1, p2 := doStuff(846601)
	fmt.Printf("846601 recipies: %s, %d\n", p1, p2)
}

func doStuff(count int) (string, int) {

	results := []byte{3, 7}

	next1 := 0
	next2 := 1
	searchstr := []byte{}
	s := strconv.Itoa(count)
	for i := 0; i < len(s); i++ {
		searchstr = append(searchstr, s[i]-'0')
	}
	slen := len(searchstr)

	found := 0
	for i := 0; i < count+12 || found == 0; i++ {
		x := results[next1] + results[next2]
		if x >= 10 {
			results = append(results, 1)
		}
		results = append(results, x%10)

		next1 = (next1 + 1 + int(results[next1])) % len(results)
		next2 = (next2 + 1 + int(results[next2])) % len(results)

		if found == 0 && len(results) > slen {
			tmp := results[len(results)-slen:]
			if reflect.DeepEqual(searchstr, tmp) {
				found = len(results) - len(searchstr)
			}
			if x >= 10 {
				tmp = results[len(results)-slen-1 : len(results)-1]
				if reflect.DeepEqual(searchstr, tmp) {
					found = len(results) - len(searchstr) - 1
				}
			}
		}

	}
	score := ""
	for i := count; i < count+10; i++ {
		score += strconv.Itoa(int(results[i]))
	}

	return score, found
}
