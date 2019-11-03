package main

import (
	"flag"
	"fmt"
	"github.com/AlexLitvak/go-quiz-game/bank"
	"os"
	"time"
)

func main() {
	filePtr := flag.String("filename", "problems.csv", "the file to process")
	limitPtr := flag.Int("limit", 2, "the timer limit")

	flag.Parse()

	timer := time.NewTimer(time.Duration(*limitPtr) * time.Second)
	correct := 0
	go func() {
		<-timer.C
		fmt.Printf("Time is up! You answered %v correct answerrs", correct)
		os.Exit(1)
	}()

	quizBank := bank.NewBank(*filePtr)
	fmt.Printf("timer limit is: %v\n", *limitPtr)

	for _, question := range quizBank.Questions {
		//fmt.Println(question)
		fmt.Printf("Question: %v \n", question.Question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if question.Answer == answer {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Printf("Oops! The correct answer is %v. Your answer was %v \n", question.Answer, answer)
		}
	}
	fmt.Printf("You scored %v out of %v. Bye bye... \n", correct, len(quizBank.Questions))

	return
}
