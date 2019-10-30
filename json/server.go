package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	// HACK: interfaceを型として保存できる。
	store  PlayerStore
	router *http.ServeMux
}

// It's quite odd (and inefficient) to be setting up a router as a request comes in and then calling it.
// リクエストのたびに、ルーティング設定をいちから作るのは不毛なので、事前に作るメソッドを用意する
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store,
		http.NewServeMux(),
	}
	// HACK: routes機能があるのか。
	// Go has a built-in routing mechanism called ServeMux (request multiplexer) which lets you attach http.Handlers to particular request paths.
	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	p.router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	return p
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
