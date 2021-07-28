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

func main() {
	csvFileName := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
	timeFlag := flag.Int("timer", timerCount, "timer for answering to a question")
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
	for i, p := range problems {
		timer := time.NewTimer(time.Duration(*timeFlag) * time.Second)
		go func() {
			select {
			case <-timer.C:
				printResults(len(problems))
				exit("Time expired.\n")
			}
		}()
		if checkAnswers(i, p.q, p.a) {
			correct++
		}
		timer.Stop()
	}
	printResults(len(problems))
}

func printResults(problems int) {
	fmt.Printf("Final score %d out of %d.", correct, problems)
}

func checkAnswers(i int, q, a string) (c bool) {
	fmt.Printf("Problem %d: %s = \n", i+1, q)
	var answer string
	fmt.Scanf("%s", &answer)
	return a == answer
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
