package url

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(ctx context.Context, userID uuid.UUID, originalURL string) error {
	url, err := NewURL(userID, originalURL)
	if err != nil {
		return err
	}

	return s.repo.Create(ctx, url)
}

func (s *Service) Get(ctx context.Context, hash string) (*URL, error) {
	return s.repo.Get(ctx, hash)
}
