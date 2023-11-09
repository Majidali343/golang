package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the number of goroutines as a command-line argument.")
		os.Exit(1)
	}

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

	doneCh := make(chan struct{})

	partialResultCh := make(chan Calculation)

	go func() {
		var totalCalculation Calculation

		for i := 0; i < numGoroutines; i++ {
			partialResult := <-partialResultCh
			totalCalculation.PunctuationCount += partialResult.PunctuationCount
			totalCalculation.VowelCount += partialResult.VowelCount
			totalCalculation.WordCount += partialResult.WordCount
			totalCalculation.LineCount += partialResult.LineCount
		}

		fmt.Printf("Total details are %+v \n", totalCalculation)

		close(doneCh)
	}()

	for i := 0; i < numGoroutines; i++ {
		go counting(data[i*segmentSize:(i+1)*segmentSize], partialResultCh, doneCh)
	}

	<-doneCh

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

func counting(data []byte, resultCh chan<- Calculation, doneCh <-chan struct{}) {
	var calculation Calculation

	for _, char := range data {
		select {
		case <-doneCh:

			return
		default:
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
	}

	resultCh <- calculation
}
