package main

import (
	"fmt"
	"sort"
)

func main() {
	population := map[string]int{

		"majid":  7,
		"zubair": 10,
		"ali":    7,
		"hamza":  8,
		"hassan": 2,
	}

	keys := make([]string, 0, len(population))
	for k := range population {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(k, population[k])
	}

}
