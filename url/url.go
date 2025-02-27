package url

import (
	"context"
	"crypto/rand"
	"math/big"
	"time"

	"github.com/google/uuid"
)

type URL struct {
	ID          uuid.UUID `json:"id" bson:"_id"`
	UserID      uuid.UUID `json:"useId,omitzero" bson:"user_id"`
	OriginalURL string    `json:"originalUrl" bson:"original_url"`
	Hash        string    `json:"hash" bson:"hash"`
	CreatedAt   time.Time `json:"createdAt" bson:"created_at"`
	ExpireAt    time.Time `json:"expireAt,omitzero" bson:"expire_at,omitzero"`
}

func NewUserURL(originalURL string, userID uuid.UUID) (*URL, error) {
	hash, err := generateHash()
	if err != nil {
		return nil, err
	}

	return &URL{
		ID:          uuid.New(),
		UserID:      userID,
		Hash:        hash,
		OriginalURL: originalURL,
		CreatedAt:   time.Now().UTC(),
	}, nil
}

func NewAnonymousURL(originalURL string) (*URL, error) {
	hash, err := generateHash()
	if err != nil {
		return nil, err
	}

	return &URL{
		ID:          uuid.New(),
		Hash:        hash,
		OriginalURL: originalURL,
		CreatedAt:   time.Now().UTC(),
		ExpireAt:    time.Now().AddDate(0, 0, 3).UTC(),
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
	CreateAnonymous(ctx context.Context, originalURL string) (*URL, error)
	CreateToUser(ctx context.Context, originalURL string, userID uuid.UUID) (*URL, error)
	Get(ctx context.Context, hash string) (*URL, error)
}
