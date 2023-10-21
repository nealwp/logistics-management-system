package service

import (
	"testing"

	"github.com/nealwp/logistics-management-system/internal/domain"
)

func TestAddItem(t *testing.T) {
    service := NewItemService()

    t.Run("empty name returns error", func(t *testing.T) {
        item := &domain.Item{Name: "", UnitPrice: 1234}
        err := service.AddItem(item)

        if err == nil {
            t.Error("expected error, but got none")
        }
    })

    t.Run("valid name does not return error", func(t *testing.T) {
        item := &domain.Item{Name: "Sample Item", UnitPrice: 1234}
        err := service.AddItem(item)

        if err != nil {
            t.Errorf("Did not expected error, but got: %v", err)
        }
    })

    
    t.Run("non-zero unit price does not return error", func(t *testing.T) {
        item := &domain.Item{Name: "Sample Item", UnitPrice: 1234}
        err := service.AddItem(item)

        if err != nil {
            t.Errorf("Did not expected error, but got: %v", err)
        }
    })
    
    t.Run("zero unit price returns error", func(t *testing.T) {
        item := &domain.Item{Name: "Sample Item", UnitPrice: 0}
        err := service.AddItem(item)

        if err == nil {
            t.Error("expected error, but got none")
        }
    })
}
