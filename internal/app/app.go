package app

import (
	"elogger/internal/controller"
	"elogger/internal/domain"
	"elogger/internal/infrastructure/logger"
	"elogger/internal/server/http"
)

func Run() {

	uc := domain.New(
		logger.New(),
	)

	ctrl := controller.New(uc)

	serv := http.New(ctrl)

	serv.RunServer()
}
