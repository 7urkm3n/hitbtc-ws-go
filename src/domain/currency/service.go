package currency

import (
	"orange-currency/src/utils/errors"
)

// Repository interface
type Repository interface {
	GetBySymbol(string) (*Currency, *errors.RestErr)
	All() ([2]*Currency, *errors.RestErr)
}

// Service interface
type Service interface {
	GetBySymbol(string) (*Currency, *errors.RestErr)
	All() ([2]*Currency, *errors.RestErr)
}

type service struct {
	repository Repository
}

// NewService returns new Service
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetBySymbol(symbol string) (*Currency, *errors.RestErr) {
	currency, err := s.repository.GetBySymbol(symbol)
	if err != nil {
		return nil, err
	}
	return currency, err
}

func (s *service) All() ([2]*Currency, *errors.RestErr) {
	data, err := s.repository.All()
	if err != nil {
		return [2]*Currency{}, err
	}
	return data, err
}
