package http

import (
	"log"
	"net/http"
)

type IMock interface {
	Mock(http.ResponseWriter, *http.Request)
}

type Http struct {
	mock IMock
}

func (s *Http) RunServer() {
	http.HandleFunc("/mock", s.mock.Mock)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

func New(m IMock) *Http {
	s := Http{
		mock: m,
	}

	s.RunServer()

	return &s
}
