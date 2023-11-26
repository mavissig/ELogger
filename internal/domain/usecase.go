package domain

type Logger interface {
	Mock()
}

type UseCase struct {
	logger Logger
}

func (uc *UseCase) Mock() {
	uc.logger.Mock()
}

func New(l Logger) *UseCase {
	return &UseCase{
		logger: l,
	}
}
