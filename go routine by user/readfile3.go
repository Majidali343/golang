package main

import (
	"fmt"
	"io/ioutil"

	"sync"
	"time"
)

func main() {

	fmt.Printf("How many go routines you want to use?\n")

	var numGoroutines int

	fmt.Print("Enter the number of goroutines: ")
	_, err := fmt.Scanln(&numGoroutines)

	if err != nil || numGoroutines <= 0 {
		fmt.Println("Invalid input. Please provide a positive integer for the number of goroutines.")
		return
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

	// segments = append(segments, data[(numGoroutines-1)*segmentSize:])

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

func counting(data []byte, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes

	punctuationCount := 0
	vowelCount := 0
	wordCount := 0
	lineCount := 0

	for _, char := range data {
		if char == '\n' {
			lineCount++
		} else if char == '!' || char == '"' || char == '$' || char == '*' || char == '.' || char == '<' || char == '?' || char == '~' || char == '{' || char == '`' {
			punctuationCount++
		} else if char == 'a' || char == 'A' || char == 'e' || char == 'E' || char == 'i' || char == 'I' || char == 'O' || char == 'o' || char == 'u' || char == 'U' {
			vowelCount++
		} else if char == '\t' || char == ' ' || char == '\n' || char == '\r' {
			wordCount++
		}
	}

	fmt.Printf("Number of lines: %d\n", lineCount)
	fmt.Printf("Number of words: %d\n", wordCount)
	fmt.Printf("Number of punctuation marks: %d\n", punctuationCount)
	fmt.Printf("Number of vowels: %d\n", vowelCount)
}
