package usecase_test

import (
	"context"
	"testing"

	"github.com/Kbnh/url_shortener/internal/domain"
	"github.com/Kbnh/url_shortener/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		inputAlias string
		mockResult string
		mockErr    error
		expectURL  string
		expectErr  bool
	}{
		{
			name:       "success",
			inputAlias: "abc123",
			mockResult: "https://google.com",
			mockErr:    nil,
			expectURL:  "https://google.com",
			expectErr:  false,
		},
		{
			name:       "not found",
			inputAlias: "nonexistent",
			mockResult: "",
			mockErr:    domain.ErrURLNotFound,
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repo := &mockRepo{
				getURLResult: tt.mockResult,
				err:          tt.mockErr,
			}
			uc := usecase.New(repo)

			url, err := uc.GetURL(context.Background(), tt.inputAlias)

			if tt.expectErr {
				assert.Error(t, err)
				if tt.mockErr != nil {
					assert.ErrorIs(t, err, tt.mockErr)
				}
				assert.Empty(t, url)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectURL, url)
			}
		})
	}
}
