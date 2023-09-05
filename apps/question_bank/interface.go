package question

import (
	"context"
	"restful-api-demo/apps/question_bank/model"
)

type QuestionInterface interface {
	SaveQuestion(context.Context, *model.Question) (*model.Question, error)
	QueryQuestion(context.Context, *model.QueryQuestionListRequest) (*model.QuestionSet, error)
	DescribeQuestion(context.Context, *model.DescribeQuestionRequest) (*model.Question, error)
	DeleteQuestion(context.Context, *model.DeleteQuestionRequest) (*model.Question, error)
	UpdateQuestion(context.Context, *model.UpdateQuestionRequest) (*model.Question, error)
}
