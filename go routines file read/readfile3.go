package main

import (
	"fmt"
	"io/ioutil"
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

	counting(data)

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Milliseconds()
	fmt.Printf("Elapsed time: %d ms\n", elapsedTime)
}

func counting(data []byte) {
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
