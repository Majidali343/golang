package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	filePath := "file.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split data into three parts
	partSize := len(data) / 3
	part1 := data[:partSize]
	part2 := data[partSize : 2*partSize]
	part3 := data[2*partSize:]

	var wg sync.WaitGroup

	wg.Add(3)
	go counting(part1, &wg)
	go counting(part2, &wg)
	go counting(part3, &wg)

	wg.Wait()

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Milliseconds()
	fmt.Printf("Elapsed time: %d ms\n", elapsedTime)
}

func counting(data []byte, wg *sync.WaitGroup) {
	defer wg.Done()
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
