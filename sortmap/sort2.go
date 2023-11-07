package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

func main() {
	population := map[string]int{
		"majid":  7,
		"zubair": 10,
		"ali":    7,
		"hamza":  8,
		"hassan": 2,
	}

	var pairs []Pair
	for k, v := range population {
		pairs = append(pairs, Pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	for _, p := range pairs {
		fmt.Printf("%v\t%v\n", p.Key, p.Value)
	}
}
