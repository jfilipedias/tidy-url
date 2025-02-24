package url

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestServiceGet(t *testing.T) {
	u, err := NewURL(uuid.New(), "https://example.com")
	if err != nil {
		t.Errorf("failed to create a URL: %v", err)
	}

	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)
	repo.
		EXPECT().
		Get(gomock.Eq(u.Hash)).
		Return(u, nil)

	s := NewService(repo)
	got, err := s.Get(u.Hash)

	assert.Nil(t, err)
	assert.Equal(t, u, got)
}
