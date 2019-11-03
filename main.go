package main

import (
	"fmt"
	"github.com/AlexLitvak/go-quiz-game/bank"
)

func main() {
	quizBank := bank.NewBank("problems.csv")
	correct := 0
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
