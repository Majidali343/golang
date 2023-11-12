// cmd/main.go
package main

import (
	"fmt"
	// "strconv"
	"time"
	"wordcount/internal/calculation"
	"wordcount/internal/file"
	"wordcount/package/counting"
	"wordcount/cmd/commands"
	
	
)




func main() {
	
	cobraargs.Execute()

	startTime := time.Now()


	data, err := file.ReadFile("../assets/" + cobraargs.FileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	segmentSize := len(data) / cobraargs.Routines

	doneCh := make(chan struct{})

	partialResultCh := make(chan calculation.Calculation)

	go func() {
		var totalCalculation calculation.Calculation

		for i := 0; i < cobraargs.Routines; i++ {
			partialResult := <-partialResultCh
			totalCalculation.PunctuationCount += partialResult.PunctuationCount
			totalCalculation.VowelCount += partialResult.VowelCount
			totalCalculation.WordCount += partialResult.WordCount
			totalCalculation.LineCount += partialResult.LineCount
		}

		fmt.Printf("Total details are %+v \n", totalCalculation)

		close(doneCh)
	}()

	for i := 0; i < cobraargs.Routines; i++ {
		go counting.Count(data[i*segmentSize:(i+1)*segmentSize], partialResultCh, doneCh)
	}

	<-doneCh

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Milliseconds()
	fmt.Printf("Elapsed time: %d ms\n", elapsedTime)
}
