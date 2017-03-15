/*
Package godux implements a state management for your backend application.
It's inspired in Redux, but with simplest concepts.
- State:   Your application state don't change.
- Actions: Your action is used in reducers, to return new value based on State.
- Reducers: Actions describe the fact that something happened, but don’t specify how the application’s state changes in response. This is the job of a reducer.

This library was inspired in Redux.
*/
package core

import "log"
import "sync"

// Store Your central store that has your application state
type Store struct {
	state       *State
	subscribers []Subscriber
	lock        sync.RWMutex
	reducer     func(*Action, *State) *State
}

type State struct {
	CurrentPerson *Person
	CurrentIndex  int
	IsFetching    bool
}

type Subscriber interface {
	Update(*State)
}

type Action struct {
	actionType int64
	value      interface{}
}

// Dispatch trigger your action type
func (store *Store) Dispatch(action *Action) {
	store.lock.RLock()
	if store.reducer == nil {
		store.lock.RUnlock()
		panic("reducer not initialized")
	}
	store.state = store.reducer(action, store.state)
	log.Printf("Dispatched action %d -> %v", action.actionType, store.state)
	for _, subscriber := range store.subscribers {
		subscriber.Update(store.state)
	}
	store.lock.RUnlock()
}

func (store *Store) Subscribe(subscriber Subscriber) {
	store.lock.RLock()
	store.subscribers = append(store.subscribers, subscriber)
	log.Printf("Subscribed %v, %d subscribers now", subscriber, len(store.subscribers))
	subscriber.Update(store.state)
	store.lock.RUnlock()
}

func (store *Store) Unsubscribe(subscriber Subscriber) {
	store.lock.RLock()
	store.subscribers = store.subscribers[:0]
	for _, existing := range store.subscribers {
		if existing != subscriber {
			store.subscribers = append(store.subscribers, existing)
		}
	}
	log.Printf("Unsubscribed %v, %d subscribers now", subscriber, len(store.subscribers))
	store.lock.RUnlock()
}

// GetState return the state of your store
func (store *Store) GetState() interface{} {
	store.lock.RLock()
	state := store.state
	store.lock.RUnlock()
	return state
}
