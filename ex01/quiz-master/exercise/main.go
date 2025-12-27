package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int64("limit", 30, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle the quiz questions")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	if *shuffle {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)

		answerCh := make(chan string)
		done := make(chan bool)

		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			select {
			case answerCh <- answer:
			case <-done:
				return
			}
		}()

		select {
		case <-timer.C:
			close(done)
			fmt.Println("\nTime Up!")
			break problemLoop
		case answer := <-answerCh:
			close(done)
			trimAnswer := strings.TrimSpace(answer)
			if trimAnswer == p.a {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Printf("Incorrect. The answer was %s\n", p.a)
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	var ret []problem
	for _, line := range lines {
		if len(line) == 2 {
			ret = append(ret, problem{
				q: strings.TrimSpace(line[0]),
				a: strings.TrimSpace(line[1]),
			})
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
