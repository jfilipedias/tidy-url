package url

import (
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(userID uuid.UUID, originalURL string) error {
	url, err := NewURL(userID, originalURL)
	if err != nil {
		return err
	}

	return s.repo.Create(url)
}
