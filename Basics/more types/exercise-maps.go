package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordDict := make(map[string]int)
	wordList := strings.Split(s, " ")
	for _ , v := range wordList {
		if k, or := wordDict[v]; or {
			wordDict[v] = k+1
		} else {
			wordDict[v] = 1
		}
	}

	return wordDict
}

func main() {
	wc.Test(WordCount)
}
