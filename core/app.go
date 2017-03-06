package core

import (
	"github.com/leejarvis/swapi"
)

const (
	ActionFetchStarted = iota
	ActionFetchSuccess = iota
	ActionFetchFailed  = iota
)

type Action struct {
	Type  int32
	Value interface{}
}

type Person struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Homeworld string   `json:"homeworld"`
	Films     []string `json:"films"`
	Species   []string `json:"species"`
	Vehicles  []string `json:"vehicles"`
	Starships []string `json:"starships"`
	Created   string   `json:"created"`
	Edited    string   `json:"edited"`
	URL       string   `json:"url"`
}

type State struct {
	CurrentPerson *Person
	CurrentIndex  int
	IsFetching    bool
}

type Subscriber interface {
	Update()
}

type Store struct {
	State      *State
	Subscriber Subscriber
}

func (store *Store) FetchNextPerson() {
	store.State = &State{
		CurrentIndex:  store.State.CurrentIndex + 1,
		CurrentPerson: store.State.CurrentPerson,
		IsFetching:    true,
	}
	store.Subscriber.Update()
	person, err := swapi.GetPerson(store.State.CurrentIndex)
	if err == nil {
		store.State = &State{
			CurrentIndex:  store.State.CurrentIndex,
			CurrentPerson: ToReyPerson(person),
			IsFetching:    false,
		}
		store.Subscriber.Update()
	} else {
		println(err)
	}
}

func (store *Store) Dispatch(action *Action) {
	switch action.Type {
	case ActionFetchFailed:
		bre
	}
	store.State = &State{
		CurrentIndex:  store.State.CurrentIndex + 1,
		CurrentPerson: store.State.CurrentPerson,
		IsFetching:    true,
	}
	store.Subscriber.Update()
	person, err := swapi.GetPerson(store.State.CurrentIndex)
	if err == nil {
		store.State = &State{
			CurrentIndex:  store.State.CurrentIndex,
			CurrentPerson: ToReyPerson(person),
			IsFetching:    false,
		}
		store.Subscriber.Update()
	} else {
		println(err)
	}
}

func ToReyPerson(person swapi.Person) *Person {
	return &Person{
		Name:      person.Name,
		Height:    person.Height,
		Mass:      person.Mass,
		HairColor: person.HairColor,
		SkinColor: person.SkinColor,
		EyeColor:  person.EyeColor,
		BirthYear: person.BirthYear,
		Gender:    person.Gender,
		Homeworld: person.Homeworld,
		Films:     person.Films,
		Species:   person.Species,
		Vehicles:  person.Vehicles,
		Starships: person.Starships,
		Created:   person.Created,
		Edited:    person.Edited,
		URL:       person.URL,
	}
}

func NewStore() *Store {
	state := &State{CurrentIndex: 0, IsFetching: false}
	return &Store{State: state}
}
