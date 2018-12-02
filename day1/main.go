package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	freq := 0
	var changes []int

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		text = strings.Replace(text, "\n","",-1)

		if text == "process" {
			process(changes)
			return
		}

		change, err := strconv.Atoi(text)
		changes = append(changes, change)

		if err != nil {
			fmt.Println(err)
			return
		}

		res := freq + change
		msg := fmt.Sprintf("Current frequency  %d, change of %d; resulting frequency  %d.",freq, change, res)
		log.Println(msg)

		freq = res

	}
}

// process solves the second part of the problem
func process(changes []int){
	m := make(map[int]int)
	m[0]= 1
	freq := 0
	for {
		for _, c := range changes {
			freq += c
			if _, ok := m[freq]; ok {
				fmt.Printf("Result: %d\n", freq)
				return
			}
			m[freq] = 1
		}
	}
}
