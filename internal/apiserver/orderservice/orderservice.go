package orderservice

import (
	"WB0/internal/apiserver/repository"
	"WB0/internal/models"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	cacheRepository *repository.CacheRepository
	sqlRepository   *repository.SqlRepository
	logger          *logrus.Logger
}

func NewOrderService(db *sql.DB, logger *logrus.Logger) *OrderService {
	return &OrderService{
		cacheRepository: repository.NewCacheRepository(),
		sqlRepository:   repository.NewSqlRepository(db),
		logger:          logger,
	}
}

func (s *OrderService) GetByUid(orderUid string) (*models.Order, error) {
	var (
		order *models.Order
		err   error
	)
	if !s.cacheRepository.IsEmpty() {
		order, err = s.cacheRepository.GetById(orderUid)
		if err != nil {
			s.logger.Info(err)
		} else {
			return order, nil
		}
	}
	order, err = s.sqlRepository.GetById(orderUid)
	if err != nil {
		s.logger.Info(err)
		return nil, err
	} else {
		s.cacheRepository.Put(orderUid, order)
		return order, nil
	}
}
