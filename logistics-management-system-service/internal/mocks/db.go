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

func (m *MockDatabase) ItemExists(itemId string) (bool, error) {
    args := m.Called(itemId)
    return args.Bool(0), args.Error(1)
}

func (m *MockDatabase) GetAllItems() ([]domain.Item, error) {
    args := m.Called()
    return args.Get(0).([]domain.Item), args.Error(1)
}
