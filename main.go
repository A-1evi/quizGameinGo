package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFileName))

	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the csv file")
	}
	problems := parseLines(lines)

	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct++
			fmt.Println("Correct!")

		}

	}
	fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for index, line := range lines {
		problems[index] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
