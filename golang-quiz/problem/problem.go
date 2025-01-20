package problem

import (
	"encoding/csv"
	"fmt"
	"os"
)

type problem struct {
	Question string
	Answer string
}


func ProblemPuller(fileName string) ([]problem, error){
	file, err := os.Open(fileName)

	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %v", fileName, err)
	}

	reader := csv.NewReader(file)

	lines, err :=reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", fileName, err)
	}

	return ProblemParser(lines) ,nil

	
} 

func ProblemParser(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{Question: line[0], Answer: line[1]}
	}

	return problems
}

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}