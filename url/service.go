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

func (s *Service) CreateAnonymous(ctx context.Context, originalURL string) (*URL, error) {
	url, err := NewAnonymousURL(originalURL)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(ctx, url)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *Service) CreateToUser(ctx context.Context, originalURL string, userID uuid.UUID) (*URL, error) {
	url, err := NewUserURL(originalURL, userID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(ctx, url)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *Service) Get(ctx context.Context, hash string) (*URL, error) {
	return s.repo.Get(ctx, hash)
}
