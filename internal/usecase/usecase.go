package usecase

type Controller interface {
	Info(string, string) error
}

type UseCase struct {
	controller Controller
}

func New(c Controller) *UseCase {
	return &UseCase{
		controller: c,
	}
}
