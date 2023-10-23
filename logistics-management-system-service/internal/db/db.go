package db

import (
	"github.com/nealwp/logistics-management-system/internal/domain"
)

type Database interface {
    InsertItem(item *domain.Item) error
    ItemExists(itemId string) (bool, error)
    InsertAllowance(allowance *domain.Allowance) error
    GetAllItems() ([]domain.Item, error)
}

