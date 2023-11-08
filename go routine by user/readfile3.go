package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {

	numGoroutines, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input for num_goroutines:", err)
		os.Exit(1)
	}

	startTime := time.Now()
	filePath := "file.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	segmentSize := len(data) / numGoroutines
	var segments [][]byte

	for i := 0; i < numGoroutines; i++ {
		segments = append(segments, data[i*segmentSize:(i+1)*segmentSize])
	}

	var wg sync.WaitGroup

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go counting(segments[i], &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Milliseconds()
	fmt.Printf("Elapsed time: %d ms\n", elapsedTime)
}

type Calculation struct {
	PunctuationCount int
	VowelCount       int
	WordCount        int
	LineCount        int
}

func counting(data []byte, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes

	var calculation Calculation

	for _, char := range data {
		if char == '\n' {
			calculation.LineCount++
		} else if char == '!' || char == '"' || char == '$' || char == '*' || char == '.' || char == '<' || char == '?' || char == '~' || char == '{' || char == '`' {
			calculation.PunctuationCount++
		} else if char == 'a' || char == 'A' || char == 'e' || char == 'E' || char == 'i' || char == 'I' || char == 'O' || char == 'o' || char == 'u' || char == 'U' {
			calculation.VowelCount++
		} else if char == '\t' || char == ' ' || char == '\n' || char == '\r' {
			calculation.WordCount++
		}
	}

	fmt.Printf("Details are %+v \n", calculation)
}
