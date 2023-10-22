package mocks

import (
    "github.com/nealwp/logistics-management-system/internal/domain"
    "github.com/stretchr/testify/mock"
)

type MockDatabase struct {
    mock.Mock
}

func (m *MockDatabase) InsertItem(item *domain.Item) error {
    args := m.Called(item)
    return args.Error(0)
}

func (m *MockDatabase) InsertAllowance(allowance *domain.Allowance) error {
    args := m.Called(allowance)
    return args.Error(0)
}

