package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filenamePtr := flag.String("filename", "problems.csv", "Path to the CSV file")
	timerInSecs := flag.Int64("time", 30, "Time in seconds")
	// Parse the command-line flags
	flag.Parse()

	file, err := os.Open(*filenamePtr) // For read access.
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	r := csv.NewReader(file)

	records, readerErr := r.ReadAll()

	if readerErr != nil {
		fmt.Println(readerErr)
	}

	runQuiz(records, *timerInSecs)
}

func runQuiz(records [][]string, time int64) {
	reader := bufio.NewReader(os.Stdin)
	correctAnswers := 0
	totalQuestions := 0

	for i := 0; i < len(records); i++ {
		if len(records[i]) != 2 {
			fmt.Println("Invalid record in CSV file:", records[i])
			continue
		}

		totalQuestions++

		fmt.Printf("Hello, what is %s?\n", records[i][0])
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid response")
			continue
		}
		answer = strings.TrimSpace(answer)

		if strings.Compare(answer, records[i][1]) == 0 {
			fmt.Println("Well done. Good Job.")
			correctAnswers++
		} else {
			fmt.Println("Sorry, Incorrect.")
		}
	}

	fmt.Printf("You got %d out of %d questions correct.\n", correctAnswers, totalQuestions)
}
