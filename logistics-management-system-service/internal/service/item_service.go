package service

import (
	"errors"
	"strings"

	"github.com/nealwp/logistics-management-system/internal/domain"
)

type ItemService struct {
    db Database
}

type Database interface {
    InsertItem(item *domain.Item) error
}

func NewItemService(db Database) *ItemService {
    return &ItemService{db: db}
}

func (s *ItemService) AddItem(item *domain.Item) error {
    
    if strings.TrimSpace(item.ID) == "" {
        return errors.New("item ID is required")
    }

    if strings.TrimSpace(item.Name) == "" {
        return errors.New("item name is required")
    }

    if item.UnitPrice <= 0 {
        return errors.New("unit price must be > 0")
    }

    s.db.InsertItem(item)

    return nil
}
