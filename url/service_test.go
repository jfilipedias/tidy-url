package url

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jfilipedias/tidy-url/constant"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var userID = uuid.New()
var originalURL = "http://example.com"

func TestServiceCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)
	repo.
		EXPECT().
		Create(gomock.Any()).
		Return(nil)

	s := NewService(repo)
	err := s.Create(userID, originalURL)

	assert.NoError(t, err)
}

func TestServiceGet(t *testing.T) {
	t.Run("existing url", func(t *testing.T) {
		u, err := NewURL(userID, originalURL)
		if err != nil {
			t.Fatalf("failed to create a URL: %v", err)
		}

		ctrl := gomock.NewController(t)
		repo := NewMockRepository(ctrl)
		repo.
			EXPECT().
			Get(gomock.Eq(u.Hash)).
			Return(u, nil)

		s := NewService(repo)
		got, err := s.Get(u.Hash)

		assert.NoError(t, err)
		assert.Equal(t, u, got)
	})

	t.Run("non-existing url", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := NewMockRepository(ctrl)
		repo.
			EXPECT().
			Get("abcdefgh").
			Return(nil, constant.ErrEntityNotFound)

		s := NewService(repo)
		u, err := s.Get("abcdefgh")

		assert.Error(t, err)
		assert.Nil(t, u)
	})
}
