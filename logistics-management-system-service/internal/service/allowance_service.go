package service

import (
	"errors"
	"strings"

	"github.com/nealwp/logistics-management-system/internal/db"
	"github.com/nealwp/logistics-management-system/internal/domain"
)

type AllowanceService struct {
    db db.Database
}

func NewAllowanceService(db db.Database) *AllowanceService {
    return &AllowanceService{db: db}
}

func (s *AllowanceService) AddAllowance(itemId string, quantity int32) error {

    if strings.TrimSpace(itemId) == "" {
        return errors.New("item id is required")
    }

    if quantity < 0 {
        return errors.New("allowance quantity >= 0")
    }

    exists, err := s.db.ItemExists(itemId)

    if err != nil {
        return err 
    }

    if !exists {
        return errors.New("item does not exist") 
    }

    allowance := &domain.Allowance{ItemId: itemId, Quantity: quantity}

    err = s.db.InsertAllowance(allowance)

    if err != nil {
        return err 
    }

    return nil
}
