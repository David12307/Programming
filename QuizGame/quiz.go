package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 5, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Scan the file and get the questions with the answers
	var questions [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")
		questions = append(questions, values)
	}

	// Ask the questions
	var correctQuestions int
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemloop:
	for i, question := range questions {
		fmt.Printf("#%d: %s = ", i+1, question[0])
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == question[1] {
				correctQuestions++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctQuestions, len(questions))
}
