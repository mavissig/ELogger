package controller

import (
	"log"
	"net/http"
)

type UseCase interface {
	Mock()
}

type Controller struct {
	usecase UseCase
}

func (c *Controller) Mock(_ http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)
	c.usecase.Mock()
}

func New(uc UseCase) *Controller {
	return &Controller{
		usecase: uc,
	}
}
