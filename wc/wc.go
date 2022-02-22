package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "I am learning Go! I am what I am"
	words := strings.Fields(sentence)
	var word_map map[string]int
	word_map = make(map[string]int)

	for _, word := range words {
		count, ok := word_map[word]
		if ok {
			count += 1
			word_map[word] = count
		} else {
			word_map[word] = 0
		}
	}
	fmt.Println(word_map)
}
