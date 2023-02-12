package orderservice

import (
	"WB0/internal/apiserver/repository"
	"WB0/internal/models"
	"context"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	cacheRepository *repository.CacheRepository
	sqlRepository   *repository.SqlRepository
	logger          *logrus.Logger
}

func NewOrderService(cacheRepository *repository.CacheRepository,
	sqlRepository *repository.SqlRepository,
	logger *logrus.Logger) *OrderService {
	return &OrderService{
		cacheRepository: cacheRepository,
		sqlRepository:   sqlRepository,
		logger:          logger,
	}
}

func (s *OrderService) GetByUid(ctx context.Context, orderUid string) (*models.Order, error) {
	var (
		order *models.Order
		err   error
	)
	if !s.cacheRepository.IsEmpty() {
		order, err = s.cacheRepository.GetById(orderUid)
		if err != nil {
			s.logger.Infof("OrderService.GetByUid error while get order from cache: %+v", err)
		} else {
			return order, nil
		}
	}
	order, err = s.sqlRepository.GetById(ctx, orderUid)
	if err != nil {
		s.logger.Infof("OrderService.GetByUid error while get order from DB: %+v", err)
		return nil, err
	} else {
		s.cacheRepository.Put(orderUid, order)
		return order, nil
	}
}

func (s *OrderService) SaveOrder(model *models.Order, ctx context.Context) error {
	err := s.sqlRepository.SaveOrder(ctx, model)
	if err != nil {
		s.logger.Infof("OrderService.SaveOrder error : %+v", err)
	}
	return err
}
