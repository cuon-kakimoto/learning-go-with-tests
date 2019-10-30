package main

import (
	"log"
	"net/http"
)

// For convenience I've made NewInMemoryPlayerStore to initialise the store, and updated the integration test to use it (store := NewInMemoryPlayerStore())
// 便利以外にあるかな？
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

// これでPlayerStoreのIFを満たす
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("cound not listen on port %v", err)
	}
}
