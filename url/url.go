package url

import (
	"time"

	"github.com/google/uuid"
	nanoid "github.com/matoous/go-nanoid/v2"
)

type URL struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	OriginalURL string
	Hash        string
	CreatedAt   time.Time
}

func NewURL(userID uuid.UUID, originalURL string) (*URL, error) {
	hash, err := generateHash()
	if err != nil {
		return nil, err
	}

	return &URL{
		ID:          uuid.New(),
		UserID:      userID,
		Hash:        hash,
		OriginalURL: originalURL,
		CreatedAt:   time.Now(),
	}, nil
}

func generateHash() (string, error) {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	return nanoid.Generate(alphabet, 54)
}

type Repository interface {
	Create(e *URL) error
	Get(hash string) (*URL, error)
}
