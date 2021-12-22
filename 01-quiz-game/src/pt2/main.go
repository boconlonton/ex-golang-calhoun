package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var scr int

	// Using `flag` standard library to process the command-line setup
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of `questions,answer`")
	quizDuration := flag.Int("time", 60, "how long does this quiz take (unit: minute")
	flag.Parse()

	// Read the CSV file
	f, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	data, err := (csv.NewReader(f)).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Load the questions
	for i, v := range data {
		fmt.Printf("Question %d: %s\n", i+1, v[0])
		ansCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansCh <- ans
		}()

		select {
		case <-time.After(time.Duration(*quizDuration) * time.Second):
			fmt.Println()
			break
		case answer := <-ansCh:
			if answer == v[1] {
				scr += 1
			}
		}

	}
	fmt.Println("--------")
	fmt.Printf("Total Score: %d pts", scr)
}
