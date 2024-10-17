package main

import "fmt"

// import (
// 	"bufio"
// 	"io"
// 	"strings"
// )

// func freq(r io.Reader) (map[string]int, error) {
// 	wordfreq := make(map[string]int)
// 	s := bufio.NewScanner(r)
// 	s.Split((bufio.ScanWords))

// 	for s.Scan() {
// 		word := strings.ToLower(s.Text())
// 		wordfreq[word]++
// 	}
// 	if err := s.Err(); err != nil {
// 		return nil, err
// 	}
// 	return wordfreq, nil
// }
func main() {
	m := make(map[string][]string)
	m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}
