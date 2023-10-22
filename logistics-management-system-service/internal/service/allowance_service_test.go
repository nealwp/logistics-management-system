package service

import (
	"errors"
	"testing"

	"github.com/nealwp/logistics-management-system/internal/mocks"
    "github.com/stretchr/testify/mock"
)

func TestAddAllowance(t *testing.T) {

    t.Run("valid allowance does not return error", func(t *testing.T) {
        mockDb := &mocks.MockDatabase{}
        mockDb.On("ItemExists", mock.Anything).Return(true, nil)
        mockDb.On("InsertAllowance", mock.Anything).Return(nil)
        service := NewAllowanceService(mockDb)
        itemId := "123456789"
        var quantity int32 = 1

        err := service.AddAllowance(itemId, quantity)

        if err != nil {
            t.Errorf("Did not expect error, but got: %v", err)
        }

        mockDb.AssertExpectations(t)
    })

    t.Run("allowance for non-existing item returns error", func(t *testing.T) {
        mockDb := &mocks.MockDatabase{}
        mockDb.On("ItemExists", mock.Anything).Return(false, nil)
        service := NewAllowanceService(mockDb)

        itemId := "111111111"
        var quantity int32 = 0

        err := service.AddAllowance(itemId, quantity)

        if err == nil {
            t.Errorf("expected error, but got none")
        }

        mockDb.AssertExpectations(t)
    })

    t.Run("empty item id returns error", func(t *testing.T) {
        mockDb := &mocks.MockDatabase{}
        service := NewAllowanceService(mockDb)

        itemId := ""
        var quantity int32 = 0

        err := service.AddAllowance(itemId, quantity)

        if err == nil {
            t.Error("expected error, but got none")
        }
    })

    t.Run("negative allowance returns error", func(t *testing.T) {
        mockDb := &mocks.MockDatabase{}
        service := NewAllowanceService(mockDb)

        itemId := "123456789"
        var quantity int32 = -1

        err := service.AddAllowance(itemId, quantity)

        if err == nil {
            t.Error("expected error, but got none")
        }
    })

    t.Run("db fail to write returns error", func(t *testing.T) {
        mockDb := &mocks.MockDatabase{}
        service := NewAllowanceService(mockDb)
        dbError := errors.New("database error")
        mockDb.On("ItemExists", mock.Anything).Return(true, nil)
        mockDb.On("InsertAllowance", mock.Anything).Return(dbError)

        itemId := "123456789"
        var quantity int32 = 1

        err := service.AddAllowance(itemId, quantity)

        if err == nil {
            t.Fatal("expected error, but got none")
        }

        if err != dbError {
            t.Fatalf("expected error %s, but got: %s", dbError, err)
        }
    })
}
