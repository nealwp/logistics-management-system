package service
 
import (
	"errors"
	"strings"

	"github.com/nealwp/logistics-management-system/internal/domain"
	"github.com/nealwp/logistics-management-system/internal/db"
)

type AllowanceService struct {
    db db.Database
}

func NewAllowanceService(db db.Database) *AllowanceService {
    return &AllowanceService{db: db}
}

func (s *AllowanceService) AddAllowance(allowance *domain.Allowance) error {
    
    if strings.TrimSpace(allowance.ID) != "" {
        return errors.New("setting allowance ID is not allowed")
    }

    if strings.TrimSpace(allowance.ItemId) == "" {
        return errors.New("item id is required")
    }

    if allowance.Quantity < 0 {
        return errors.New("allowance quantity >= 0")
    }

    err := s.db.InsertAllowance(allowance)

    if err != nil {
        return err 
    }

    return nil
}
