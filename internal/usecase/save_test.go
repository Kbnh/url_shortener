package usecase_test

import (
	"context"
	"testing"

	"github.com/Kbnh/url_shortener/internal/domain"
	"github.com/Kbnh/url_shortener/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSaveURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		inputURL   string
		inputAlias string
		mockErr    error
		expectID   int64
		expectErr  bool
	}{
		{
			name:       "empty alias test",
			inputURL:   "https://google.com",
			inputAlias: "",
			mockErr:    nil,
			expectID:   111,
			expectErr:  false,
		},
		{
			name:       "alias duplicate",
			inputURL:   "https://google.com",
			inputAlias: "duplicate",
			mockErr:    domain.ErrURLExists,
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repo := &mockRepo{err: tt.mockErr}
			uc := usecase.New(repo)

			id, err := uc.SaveURL(context.Background(), tt.inputURL, tt.inputAlias)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, int64(0), id)
				if tt.mockErr != nil {
					assert.ErrorIs(t, err, tt.mockErr)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectID, id)
			}
		})
	}
}
