package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var ans string
	var scr int

	// Using `flag` standard library to process the command-line setup
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of `questions,answer`")
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
		_, err := fmt.Scanf("%s\n", &ans)
		if err != nil {
			log.Fatal(err)
		}
		if ans == v[1] {
			scr += 1
		}
	}
	fmt.Println("--------")
	fmt.Printf("Total Score: %d pts", scr)
}
