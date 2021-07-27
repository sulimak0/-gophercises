package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csf file: %s", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file.")
	}

	problems := parseLines(lines)
	fmt.Println(problems)

	correct := 0
	for i, p := range problems {
		if checkAnswers(i, p.q, p.a) {
			correct++
		}
	}
	fmt.Printf("Final score %d out of %d.", correct, len(problems))

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
