package usecase_test

import "context"

type mockRepo struct {
	getURLResult string
	err          error
}

func (m *mockRepo) SaveURL(ctx context.Context, url, alias string) (int64, error) {
	if m.err != nil {
		return 0, m.err
	}
	return 111, nil
}

func (m *mockRepo) GetURL(ctx context.Context, alias string) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.getURLResult, nil
}

func (m *mockRepo) DeleteURL(ctx context.Context, id int64) error {
	return m.err
}
