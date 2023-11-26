package controller

import (
	"elogger/internal/domain"
	"encoding/json"
	"log"
	"net/http"
)

type UseCase interface {
	Mock()
	LogIP(domain.LogIP__)
}

type Controller struct {
	usecase UseCase
}

func (c *Controller) Mock(_ http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)
	c.usecase.Mock()
}

func (c *Controller) LogIP(_ http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("log")
	if err != nil {
		log.Println(err)
		return
	}

	l := domain.LogIP__{}

	err = json.NewDecoder(file).Decode(&l)
	if err != nil {
		log.Println(err)
		return
	}
	c.usecase.LogIP(l)
}

func New(uc UseCase) *Controller {
	return &Controller{
		usecase: uc,
	}
}
