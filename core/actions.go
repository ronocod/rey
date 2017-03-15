package core

import (
	"github.com/leejarvis/swapi"
)

const (
	ActionFetchPersonRequest = iota
	ActionFetchPersonSuccess = iota
	ActionFetchPersonFailed  = iota
)

func newAction(actionType int64) *Action {
	return &Action{actionType: actionType}
}

func FetchNextPersonRequestAction() *Action {
	return newAction(ActionFetchPersonRequest)
}

func FetchNextPersonSuccessAction(person swapi.Person) *Action {
	return newActionWithValue(ActionFetchPersonSuccess, person)
}

func FetchNextPersonFailedAction() *Action {
	return newAction(ActionFetchPersonFailed)
}

func newActionWithValue(actionType int64, value interface{}) *Action {
	return &Action{actionType: actionType, value: value}
}
