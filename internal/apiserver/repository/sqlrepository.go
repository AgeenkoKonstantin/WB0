package repository

import (
	"WB0/internal/models"
	"database/sql"
)

type SqlRepository struct {
	db *sql.DB
}

func NewSqlRepository(db *sql.DB) *SqlRepository {
	return &SqlRepository{
		db: db,
	}
}

func (repo *SqlRepository) GetById(orderUID string) (*models.Order, error) {
	var ord models.Order

	if err := repo.db.QueryRow(getOrderByOrderUIDQuery, orderUID).Scan(
		&ord.OrderUID,
		&ord.TrackNumber,
		&ord.Entry,
		&ord.Locale,
		&ord.InternalSignature,
		&ord.CustomerId,
		&ord.DeliveryService,
		&ord.Shardkey,
		&ord.SmId,
		&ord.DateCreated,
		&ord.OofShard,
	); err != nil {
		return nil, err
	}

	if err := repo.db.QueryRow(getDeliveryByOrderUIDQuery, orderUID).Scan(
		&ord.Delivery.Name,
		&ord.Delivery.Phone,
		&ord.Delivery.Zip,
		&ord.Delivery.City,
		&ord.Delivery.Address,
		&ord.Delivery.Region,
		&ord.Delivery.Email,
	); err != nil {
		return nil, err
	}

	if err := repo.db.QueryRow(getPaymentByOrderUIDQuery, orderUID).Scan(
		&ord.Payment.Transaction,
		&ord.Payment.RequestId,
		&ord.Payment.Currency,
		&ord.Payment.Provider,
		&ord.Payment.Amount,
		&ord.Payment.PaymentDt,
		&ord.Payment.Bank,
		&ord.Payment.DeliveryCost,
		&ord.Payment.GoodsTotal,
		&ord.Payment.CustomFee,
	); err != nil {
		return nil, err
	}

	rows, err := repo.db.Query(getItemsByOrderUIDQuery, orderUID)
	if err != nil {
		return nil, err
	}

	item := models.Item{}

	for rows.Next() {
		err := rows.Scan(
			&item.ChrtId,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmId,
			&item.Brand,
			&item.Status,
		)
		if err != nil {
			return nil, err
		}
		ord.Items = append(ord.Items, item)
	}

	return &ord, nil
}
