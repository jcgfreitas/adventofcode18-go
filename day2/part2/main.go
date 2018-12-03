package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main(){
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(data)
	ids := strings.Split(text, "\n")

	g := make([][]int, len(ids))
	for i := range g {
		g[i] = make([]int, len(ids))
	}

	for i := 0 ; i < len(ids[0]) ; i ++ {
		m := make(map[uint8][]int)
		for j, id := range ids{
			c := id[i]
			s := m[c]
			for _, pos := range s{
				g[pos][j] +=1
				if g[pos][j] == len(ids[0]) -1 {
					fmt.Printf("result: %s\n", result(ids[pos], ids[j]))
				}
			}
			m[c] = append(s, j)
		}
	}
}

func result(s1 string, s2 string) string {
	var res []uint8
	for i :=0;i < len(s1);i++{
		if s1[i] == s2 [i]{
			res = append(res, s1[i])
		}
	}
	return string(res)
}
