package domain

type Logger interface {
	Mock()
	LogIP(LogIP__)
}

type UseCase struct {
	logger Logger
}

func (uc *UseCase) Mock() {
	uc.logger.Mock()
}

func (uc *UseCase) LogIP(ip__ LogIP__) {
	uc.logger.LogIP(ip__)
}

func New(l Logger) *UseCase {
	return &UseCase{
		logger: l,
	}
}
