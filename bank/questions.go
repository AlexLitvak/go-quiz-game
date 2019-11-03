package bank

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type QuestionsBank struct {
	Questions []QuickQuestion
}

func NewBank(filename string) *QuestionsBank {
	b := QuestionsBank{Questions: nil}
	b.loadQuestions(filename)
	return &b
}

type QuickQuestion struct {
	Question string
	Answer   string
}

func (q QuickQuestion) String() string {
	return fmt.Sprintf("Question: %v, Answer: %v", q.Question, q.Answer)
}

func (bank *QuestionsBank) loadQuestions(filename string) {
	bank.Questions = make([]QuickQuestion, 0)
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Could not open csv file", err)
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("error while reading ", err)
		}

		bank.Questions = append(bank.Questions, QuickQuestion{Question: record[0], Answer: strings.TrimSpace(record[1])})
	}

}
