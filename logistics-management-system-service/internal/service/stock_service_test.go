package service

import (
	"errors"
	"testing"

	"github.com/nealwp/logistics-management-system/internal/domain"
	"github.com/nealwp/logistics-management-system/internal/mocks"
    "github.com/stretchr/testify/mock"
)

func TestCreateStockRecord(t *testing.T) {

    t.Run("valid item does not return error", func(t *testing.T) {
        mockDb := &mocks.MockDatabase{}
        mockDb.On("InsertItem", mock.Anything).Return(nil)
        service := NewItemService(mockDb)
        item := &domain.Item{ID: "123456789", Name: "Sample Item", UnitPrice: 1234}

        err := service.AddItem(item)

        if err != nil {
            t.Errorf("Did not expect error, but got: %v", err)
        }

        mockDb.AssertExpectations(t)
    })


    t.Run("empty id returns error", func(t *testing.T) {
        mockDb := &mocks.MockDatabase{}
        service := NewItemService(mockDb)
        item := &domain.Item{ID: "", Name: "Sample Item", UnitPrice: 0}

        err := service.AddItem(item)

        if err == nil {
            t.Error("expected error, but got none")
        }
    })

    t.Run("db fail to write returns error", func(t *testing.T) {
        mockDb := &mocks.MockDatabase{}
        service := NewItemService(mockDb)
        item := &domain.Item{ID: "123456789", Name: "Sample Item", UnitPrice: 1234}

        dbError := errors.New("database error")
        mockDb.On("InsertItem", mock.Anything).Return(dbError)

        err := service.AddItem(item)

        if err == nil {
            t.Fatal("expected error, but got none")
        }

        if err != dbError {
            t.Fatalf("expected error %s, but got: %s", dbError, err)
        }
    })
}
