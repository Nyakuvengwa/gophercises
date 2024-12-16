package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World. Lets do some math")

	file, err := os.Open("problems.csv") // For read access.
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	r := csv.NewReader(file)

	records, readerErr := r.ReadAll()

	if readerErr != nil {
		fmt.Println(readerErr)
	}

	runQuiz(records)
}

func runQuiz(records [][]string) {
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

		if answer == records[i][1] {
			fmt.Println("Well done. Good Job.")
			correctAnswers++
		} else {
			fmt.Println("Sorry, Incorrect.")
		}
	}

	fmt.Printf("You got %d out of %d questions correct.\n", correctAnswers, totalQuestions)
}
