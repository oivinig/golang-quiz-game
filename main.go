package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type CSVProblem struct {
	Question string
	Answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "set time limit to answer questions")
	flag.Parse()

	answeerSheet := loadData(*csvFileName)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var correct int
	for i, p := range answeerSheet {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.Question)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d. \n", correct, len(answeerSheet))
			return
		case answer := <-answerCh:
			if answer == p.Answer {
				correct++
			}
		}
		fmt.Printf("You scored %d out of %d. \n", correct, len(answeerSheet))
	}

}

func loadData(filename string) []CSVProblem {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open the CSV file.")
		os.Exit(1)
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the provided CSV file.")
		os.Exit(1)
	}
	defer file.Close()

	var answerSheet []CSVProblem
	for _, record := range records {
		data := CSVProblem{
			Question: record[0],
			Answer:   record[1],
		}
		answerSheet = append(answerSheet, data)
	}
	return answerSheet
}
