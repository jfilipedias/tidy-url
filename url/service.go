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

func (s *Service) CreateAnonymous(ctx context.Context, originalURL string) error {
	url, err := NewAnonymousURL(originalURL)
	if err != nil {
		return err
	}

	return s.repo.Create(ctx, url)
}

func (s *Service) CreateToUser(ctx context.Context, originalURL string, userID uuid.UUID) error {
	url, err := NewUserURL(originalURL, userID)
	if err != nil {
		return err
	}

	return s.repo.Create(ctx, url)
}

func (s *Service) Get(ctx context.Context, hash string) (*URL, error) {
	return s.repo.Get(ctx, hash)
}
