package http

import (
	"log"
	"net/http"
)

type Controller interface {
	Mock(http.ResponseWriter, *http.Request)
	LogIP(http.ResponseWriter, *http.Request)
}

type Http struct {
	controller Controller
}

func (s *Http) registry() {
	http.HandleFunc("/mock", s.controller.Mock)
	http.HandleFunc("/info", s.controller.LogIP)
}

func (s *Http) RunServer() {
	s.registry()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

func New(c Controller) *Http {
	s := Http{
		controller: c,
	}
	return &s
}
