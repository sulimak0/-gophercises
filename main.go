package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var timerCount = 5
var correct = 0
var isRandomized bool
var timeFlag *int

type problem struct {
	question string
	answer string
}

func main() {
	csvFileName := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
	timeFlag = flag.Int("timer", timerCount, "timer for answering to a question")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csf file: %s", *csvFileName))
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file.")
	}
	problems := parseLines(lines)

	fmt.Println("Please press Enter key to start the quiz:")
	fmt.Scanln()

	processAnswers(problems)

	printResults(len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func processAnswers(problems []problem) {
	for i, problem := range problems {
		timer := time.NewTimer(time.Duration(*timeFlag) * time.Second)
		go func() {
			select {
			case <-timer.C:
				printResults(len(problems))
				exit("Time expired.")
			}
		}()
		if checkAnswers(i+1, problem.question, problem.answer) {
			correct++
		}
		timer.Stop()
	}
}

func checkAnswers(i int, question, answer string) bool {
	fmt.Printf("Problem %d: %s = \n", i, question)
	var userAnswer string
	fmt.Scanf("%s", &userAnswer)
	return answer == userAnswer
}

func printResults(problems int) {
	fmt.Printf("Final score %d out of %d.\n", correct, problems)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
