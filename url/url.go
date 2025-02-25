package url

import (
	"context"
	"crypto/rand"
	"math/big"
	"time"

	"github.com/google/uuid"
)

type URL struct {
	ID          uuid.UUID `bson:"_id"`
	UserID      uuid.UUID `bson:"user_id"`
	OriginalURL string    `bson:"original_url"`
	Hash        string    `bson:"hash"`
	CreatedAt   time.Time `bson:"created_at"`
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

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const length = 8

func generateHash() (string, error) {
	var result []byte
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result = append(result, charset[num.Int64()])
	}
	return string(result), nil
}

type Repository interface {
	Create(ctx context.Context, u *URL) error
	Get(ctx context.Context, hash string) (*URL, error)
}

type UseCase interface {
	Create(ctx context.Context, userID uuid.UUID, originalURL string) error
	Get(ctx context.Context, hash string) (*URL, error)
}
