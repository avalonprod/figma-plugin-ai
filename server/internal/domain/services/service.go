package services

type Service struct {
}

type Deps struct {
}

func NewService(deps *Deps) *Service {
	return &Service{}
}
