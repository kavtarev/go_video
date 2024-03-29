package main

import "net/http"

type ServerApi struct {
	storage Storage
	address string
}

func NewServerApi(storage Storage, address string) *ServerApi {
	return &ServerApi{storage: storage, address: address}
}

func (s *ServerApi) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user", s.GetUser)

	http.ListenAndServe(s.address, mux)
}

func (s *ServerApi) GetUser(w http.ResponseWriter, r *http.Request) {
	 s.storage.getUser()
}