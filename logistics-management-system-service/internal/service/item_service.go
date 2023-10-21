package service

import (
	"errors"
	"strings"

	"github.com/nealwp/logistics-management-system/internal/domain"
)

type ItemService struct {}

func NewItemService() *ItemService {
    return &ItemService{}
}

func (s *ItemService) AddItem(item *domain.Item) error {
    if strings.TrimSpace(item.Name) == "" {
        return errors.New("item name is required")
    }

    if item.UnitPrice <= 0 {
        return errors.New("unit price must be > 0")
    }

    return nil
}
