package core

import (
	"github.com/leejarvis/swapi"
)

const (
	FetchPersonRequest ActionType = iota
	FetchPersonSuccess
	FetchPersonFailed
)

func newAction(actionType ActionType) *Action {
	return &Action{actionType: actionType}
}

func FetchNextPersonRequestAction() *Action {
	return newAction(FetchPersonRequest)
}

func FetchNextPersonSuccessAction(person swapi.Person) *Action {
	return newActionWithValue(FetchPersonSuccess, person)
}

func FetchNextPersonFailedAction() *Action {
	return newAction(FetchPersonFailed)
}

func newActionWithValue(actionType ActionType, value interface{}) *Action {
	return &Action{actionType: actionType, value: value}
}
