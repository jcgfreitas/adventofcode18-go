package main

import (
	"fmt"
	"io/ioutil"
	"unicode/utf8"
)

func main (){
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var count2, count3, temp2, temp3 int

	h := make(map[rune]int)

	for i, w := 0, 0; i < len(data); i += w {
		r, width := utf8.DecodeRune(data[i:])
		w = width

		if r == '\n'{
			for _, v := range h{
				if v == 2 {
					temp2 = 1
				}
				if v == 3 {
					temp3 = 1
				}
			}
			count2 += temp2
			count3 += temp3
			temp2 = 0
			temp3 = 0
			h = make(map[rune]int)
			continue
		}
		h[r] += 1
	}

	fmt.Printf("Result: %d * %d = %d\n", count2, count3, count2*count3)

}
