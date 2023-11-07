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

	lines := countLines(data)
	words := countWords(data)
	punctuation := countPunctuation(data)
	vowels := countVowels(data)

	fmt.Printf("Number of lines: %d\n", lines)
	fmt.Printf("Number of words: %d\n", words)
	fmt.Printf("Number of punctuation marks: %d\n", punctuation)
	fmt.Printf("Number of vowels: %d\n", vowels)

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Milliseconds()
	fmt.Printf("Elapsed time: %d ms\n", elapsedTime)
}

func countLines(data []byte) int {
	lineCount := 0
	for _, char := range data {
		if char == '\n' {
			lineCount++
		}
	}
	return lineCount
}

func isSpace(char byte) bool {
	spaces := []byte{'\t', ' ', '\n', '\r'}
	for _, s := range spaces {
		if char == s {
			return true
		}
	}
	return false
}

func countWords(data []byte) int {
	wordCount := 0
	inWord := false
	for _, char := range data {
		if isSpace(char) {
			inWord = false
		} else if !inWord {
			wordCount++
			inWord = true
		}
	}
	return wordCount
}



func isPunct(char byte) bool {
	punctuations := []byte{'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}
	for _, p := range punctuations {
		if char == p {
			return true
		}
	}
	return false
}

func countPunctuation(data []byte) int {
	punctuationCount := 0
	for _, char := range data {
		if isPunct(char) {
			punctuationCount++
		}
	}
	return punctuationCount
}

func isVowel(char byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, v := range vowels {
		if char == v {
			return true
		}
	}
	return false
}

func countVowels(data []byte) int {
	vowelCount := 0
	for _, char := range data {
		if isVowel(char) {
			vowelCount++
		}
	}
	return vowelCount
}
