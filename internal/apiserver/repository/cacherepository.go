package repository

import (
	"WB0/internal/cache"
	"WB0/internal/models"
	"encoding/json"
)

type CacheRepository struct {
	cache *cache.Cache
}

func NewCacheRepository() *CacheRepository {
	var ch cache.Cache
	ch.Data = make(map[string]string)

	return &CacheRepository{
		cache: &ch,
	}
}

func (repo *CacheRepository) GetById(uid string) (*models.Order, error) {
	result, err := repo.cache.Get(uid)
	if err != nil {
		return nil, err
	}
	var res models.Order
	if err := json.Unmarshal([]byte(result), &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *CacheRepository) Put(uid string, order *models.Order) error {
	orderBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}
	repo.cache.Put(order.OrderUID, string(orderBytes))
	return nil
}

func (repo *CacheRepository) IsEmpty() bool {
	return len(repo.cache.Data) == 0
}
