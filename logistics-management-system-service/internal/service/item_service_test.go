package service

import (
	"testing"

	"github.com/nealwp/logistics-management-system/internal/domain"
)

func TestAddItem(t *testing.T) {
    service := NewItemService()

    t.Run("empty name returns error", func(t *testing.T) {
        item := &domain.Item{Name: ""}
        err := service.AddItem(item)

        if err == nil {
            t.Error("expected error, but got none")
        }
    })

    t.Run("valid name does not return error", func(t *testing.T) {
        item := &domain.Item{Name: "Sample Item"}
        err := service.AddItem(item)

        if err != nil {
            t.Error("expected error, but got none")
        }
    })
}
