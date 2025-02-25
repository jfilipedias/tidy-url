package url_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jfilipedias/tidy-url/constant"
	"github.com/jfilipedias/tidy-url/url"
	"github.com/jfilipedias/tidy-url/url/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userID = uuid.New()
var originalURL = "http://example.com"

func TestServiceCreate(t *testing.T) {
	repo := mocks.NewRepository(t)
	repo.
		On("Create", mock.Anything, mock.Anything).
		Return(nil)

	s := url.NewService(repo)
	err := s.Create(context.Background(), userID, originalURL)

	assert.NoError(t, err)
}

func TestServiceGet(t *testing.T) {
	t.Run("existing url", func(t *testing.T) {
		u, err := url.NewURL(userID, originalURL)
		if err != nil {
			t.Fatalf("failed to create a URL: %v", err)
		}

		repo := mocks.NewRepository(t)
		repo.
			On("Get", mock.Anything, u.Hash).
			Return(u, nil)

		s := url.NewService(repo)
		got, err := s.Get(context.Background(), u.Hash)

		assert.NoError(t, err)
		assert.Equal(t, u, got)
	})

	t.Run("non-existing url", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.
			On("Get", mock.Anything, mock.Anything).
			Return(nil, constant.ErrEntityNotFound)

		s := url.NewService(repo)
		u, err := s.Get(context.Background(), "abcdefgh")

		assert.Error(t, err)
		assert.Nil(t, u)
	})
}
