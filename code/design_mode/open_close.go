package main

import (
	"context"
	"fmt"
)

// 拉题-->初始化-->答题-->订正-->加积分-->主观题批改-->主观题批改回传
type Paper struct {
	ID        int
	Answerer  string
	Questions []string
}

type Examination interface {
	Execute(ctx context.Context, paper Paper)
}

type PullPaper struct {
}

func (*PullPaper) Execute(ctx context.Context, paper Paper) {
	fmt.Println("拉取试卷")
}

type PaperInit struct {
}

func (*PaperInit) Execute(ctx context.Context, paper Paper) {
	fmt.Println("试卷初始化")
}

type Answer struct {
}

func (*Answer) Execute(ctx context.Context, paper Paper) {
	fmt.Println("答题")
}

type Correct struct {
}

func (*Correct) Execute(ctx context.Context, paper Paper) {
	fmt.Println("批改")
}

func main() {
	paper := Paper{}
	exams := make([]Examination, 0)
	exams = append(exams, &PullPaper{})
	exams = append(exams, &PaperInit{})
	exams = append(exams, &Answer{})
	exams = append(exams, &Correct{})
	ctx := context.Background()
	for _, exam := range exams {
		exam.Execute(ctx, paper)
	}
}
