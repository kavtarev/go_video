package main

type Storage interface {
	getUser() (any, error)
}

type NewMockStorage struct {
	a string
}

func (s *NewMockStorage) getUser() (any, error){
	return "new user", nil
}