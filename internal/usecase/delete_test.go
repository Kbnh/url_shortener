package usecase_test

import (
	"context"
	"testing"

	"github.com/Kbnh/url_shortener/internal/domain"
	"github.com/Kbnh/url_shortener/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestDeleteURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		inputID   int64
		mockErr   error
		expectErr bool
	}{
		{
			name:      "success",
			inputID:   111,
			mockErr:   nil,
			expectErr: false,
		},
		{
			name:      "not found",
			inputID:   999,
			mockErr:   domain.ErrURLNotFound,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repo := &mockRepo{err: tt.mockErr}
			uc := usecase.New(repo)

			err := uc.DeleteURL(context.Background(), tt.inputID)

			if tt.expectErr {
				assert.Error(t, err)
				if tt.mockErr != nil {
					assert.ErrorIs(t, err, tt.mockErr)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
