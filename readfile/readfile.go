package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	startTime := time.Now()
	inWord := false
	filePath := "file.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := countLines1(data)
	words := countWords1(data, inWord)
	punctuation := countPunctuation1(data)
	vowels := countVowels1(data)

	fmt.Printf("Number of lines: %d\n", lines)
	fmt.Printf("Number of words: %d\n", words)
	fmt.Printf("Number of punctuation marks: %d\n", punctuation)
	fmt.Printf("Number of vowels: %d\n", vowels)

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Milliseconds()
	fmt.Printf("Elapsed time: %d ms\n", elapsedTime)
}

func countLines1(data []byte) int {
	lineCount := 0
	for _, char := range data {
		if char == '\n' {
			lineCount++
		}
	}
	return lineCount
}

func countWords1(data []byte, inWord bool) int {
	wordCount := 0
	spaces := []byte{'\t', ' ', '\n', '\r'}

	for _, char := range data {
		for _, s := range spaces {

			if char == s {
				inWord = false
			} else {
				wordCount++
				inWord = true
			}
		}
	}
	return wordCount
}

func countPunctuation1(data []byte) int {
	punctuationCount := 0
	punctuations := []byte{'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}
	for _, char := range data {
		for _, p := range punctuations {
			if char == p {
				punctuationCount++
			}
		}
	}
	return punctuationCount
}

func countVowels1(data []byte) int {
	vowelCount := 0
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, char := range data {
		for _, v := range vowels {
			if char == v {
				vowelCount++
			}
		}
	}
	return vowelCount
}
