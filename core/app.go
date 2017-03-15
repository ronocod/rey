package core

import (
	"github.com/leejarvis/swapi"
)

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

func reduce(action *Action, state *State) *State {
	switch action.actionType {
	case FetchPersonRequest:
		return &State{
			CurrentIndex:  state.CurrentIndex + 1,
			CurrentPerson: state.CurrentPerson,
			IsFetching:    true,
		}
	case FetchPersonSuccess:
		return &State{
			CurrentIndex:  state.CurrentIndex,
			CurrentPerson: ToReyPerson(action.value.(swapi.Person)),
			IsFetching:    false,
		}
	case FetchPersonFailed:
		return &State{
			CurrentIndex:  state.CurrentIndex,
			CurrentPerson: state.CurrentPerson,
			IsFetching:    false,
		}
	default:
		return state
	}
}

func NewStore() *Store {
	return &Store{
		state:   &State{CurrentIndex: 0, IsFetching: false},
		reducer: reduce,
	}
}

func FetchNextPerson(store *Store) {
	go fetchPerson(store.state.CurrentIndex, store)
}

func fetchPerson(index int, store *Store) {
	store.Dispatch(FetchNextPersonRequestAction())
	person, err := swapi.GetPerson(store.state.CurrentIndex)
	if err == nil {
		store.Dispatch(FetchNextPersonSuccessAction(person))
	} else {
		store.Dispatch(FetchNextPersonFailedAction())
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
