package service

import (
	"errors"
	"testing"

	"github.com/nealwp/logistics-management-system/internal/domain"
	"github.com/nealwp/logistics-management-system/internal/mocks"
    "github.com/stretchr/testify/mock"
)

func TestAddAllowance(t *testing.T) {

    t.Run("valid allowance does not return error", func(t *testing.T) {
        mockDb := &mocks.MockDatabase{}
        mockDb.On("InsertAllowance", mock.Anything).Return(nil)
        service := NewAllowanceService(mockDb)
        allowance := &domain.Allowance{ItemId: "123456789", Quantity: 1}

        err := service.AddAllowance(allowance)

        if err != nil {
            t.Errorf("Did not expect error, but got: %v", err)
        }

        mockDb.AssertExpectations(t)
    })


    t.Run("empty item id returns error", func(t *testing.T) {
        mockDb := &mocks.MockDatabase{}
        service := NewAllowanceService(mockDb)
        allowance := &domain.Allowance{ItemId: "", Quantity: 1}

        err := service.AddAllowance(allowance)

        if err == nil {
            t.Error("expected error, but got none")
        }
    })

    t.Run("db fail to write returns error", func(t *testing.T) {
        mockDb := &mocks.MockDatabase{}
        service := NewAllowanceService(mockDb)
        allowance := &domain.Allowance{ItemId: "123456789", Quantity: 1}

        dbError := errors.New("database error")
        mockDb.On("InsertAllowance", mock.Anything).Return(dbError)

        err := service.db.InsertAllowance(allowance)

        if err == nil {
            t.Fatal("expected error, but got none")
        }

        if err != dbError {
            t.Fatalf("expected error %s, but got: %s", dbError, err)
        }
    })
}
