package main

import (
	"flag"
	"fmt"
	"quiz-app/problem"
	"time"
)

func main() {
	fileName := flag.String("f", "quiz.csv", "path of CSV file containing the questions")
	timeDuration := flag.Int("t", 30, "timer or the quiz round")


	correctAns := 0

	flag.Parse()

	problems, err := problem.ProblemPuller(*fileName)

	if err != nil {
		// Exit with msg
		problem.Exit(fmt.Sprint(err))
		
	}


	answerChan := make(chan string)

	//NewTimer creates a new Timer that will send the current time on its channel after at least duration d.
	timer := time.NewTimer(time.Duration(*timeDuration) * time.Second)

	// Loop over []problem

	for i, prob := range problems {

		var userAnswer string
		fmt.Printf("Question %d: %s = ", i+1, prob.Question)

		// using go routine
		go func ()  {
			fmt.Scanf("%s", &userAnswer)
			
			answerChan <- userAnswer
		}()


		select {
			case <-timer.C:
				// if reached here means timeout occur
				fmt.Println("\nTimeout")

				fmt.Printf("You scored %d out of %d \n", correctAns, len(problems))
				return
		
			case userResponse := <-answerChan:
				if userResponse == prob.Answer {
					correctAns++
				}

				// after going through all question, closing the channel : answerChan := make(chan string)
				if i == len(problems)-1 {
					close(answerChan)
				}
		}

		

	}

	// if reached here means attempted all question
	fmt.Printf("You scored %d out of %d \n", correctAns, len(problems))

	fmt.Println("Press Enter to exit")
	<-answerChan

	


}