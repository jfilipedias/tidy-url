package url

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
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
	if err != nil {

	}
	if !reflect.DeepEqual(u, got) {
		t.Errorf("want %v, but got %v", u, got)
	}
}
