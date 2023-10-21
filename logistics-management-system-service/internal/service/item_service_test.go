package service

import (
	"testing"

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

func TestAddItem(t *testing.T) {

    t.Run("empty name returns error", func(t *testing.T) {
        mockDb := new(MockDatabase)
        service := NewItemService(mockDb)
        item := &domain.Item{ID: "123456789", Name: "", UnitPrice: 1234}

        err := service.AddItem(item)

        if err == nil {
            t.Error("expected error, but got none")
        }
    })

    t.Run("valid name does not return error", func(t *testing.T) {
        mockDb := new(MockDatabase)
        mockDb.On("InsertItem", mock.Anything).Return(nil)
        service := NewItemService(mockDb)
        item := &domain.Item{ID: "123456789", Name: "Sample Item", UnitPrice: 1234}

        err := service.AddItem(item)

        if err != nil {
            t.Errorf("Did not expected error, but got: %v", err)
        }

    })

    
    t.Run("non-zero unit price does not return error", func(t *testing.T) {
        mockDb := new(MockDatabase)
        service := NewItemService(mockDb)
        mockDb.On("InsertItem", mock.Anything).Return(nil)
        item := &domain.Item{ID: "123456789", Name: "Sample Item", UnitPrice: 1234}

        err := service.AddItem(item)

        if err != nil {
            t.Errorf("Did not expect error, but got: %v", err)
        }
    })
    
    t.Run("zero unit price returns error", func(t *testing.T) {
        mockDb := new(MockDatabase)
        service := NewItemService(mockDb)
        item := &domain.Item{ID: "123456789", Name: "Sample Item", UnitPrice: 0}

        err := service.AddItem(item)

        if err == nil {
            t.Error("expected error, but got none")
        }
    })
}
