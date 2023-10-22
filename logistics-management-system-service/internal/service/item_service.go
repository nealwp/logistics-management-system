package service

import (
	"errors"
	"strings"

	"github.com/nealwp/logistics-management-system/internal/domain"
	"github.com/nealwp/logistics-management-system/internal/db"
)

type ItemService struct {
    db db.Database
}

func NewItemService(db db.Database) *ItemService {
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

    err := s.db.InsertItem(item)

    if err != nil {
        return err 
    }

    return nil
}

func (s *ItemService) ItemExists(itemId string) (bool, error) {
    
    exists, err := s.db.ItemExists(itemId)

    if err != nil {
        return false, err 
    }

    return exists, nil
}
