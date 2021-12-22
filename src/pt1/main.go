package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	var ans string
	var scr int

	fn := os.Args[1]
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	data, err := (csv.NewReader(f)).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
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
