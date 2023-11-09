// cmd/main.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"wordcount/internal/calculation"
	"wordcount/internal/file"
	"wordcount/package/counting"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the number of goroutines as a command-line argument.")
		os.Exit(1)
	}

	numGoroutines, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input for num_goroutines:", err)
		os.Exit(1)
	}

	startTime := time.Now()
	filePath := os.Args[2]
	data, err := file.ReadFile("../assets/" + filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	segmentSize := len(data) / numGoroutines

	doneCh := make(chan struct{})

	partialResultCh := make(chan calculation.Calculation)

	go func() {
		var totalCalculation calculation.Calculation

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
		go counting.Count(data[i*segmentSize:(i+1)*segmentSize], partialResultCh, doneCh)
	}

	<-doneCh

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Milliseconds()
	fmt.Printf("Elapsed time: %d ms\n", elapsedTime)
}
