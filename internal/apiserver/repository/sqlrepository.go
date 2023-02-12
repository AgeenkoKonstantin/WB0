package repository

import (
	"WB0/internal/models"
	"context"
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

func (repo *SqlRepository) GetById(ctx context.Context, orderUID string) (*models.Order, error) {
	var ord models.Order

	if err := repo.db.QueryRowContext(ctx, getOrderByOrderUIDQuery, orderUID).Scan(
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

	if err := repo.db.QueryRowContext(ctx, getDeliveryByOrderUIDQuery, orderUID).Scan(
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

	if err := repo.db.QueryRowContext(ctx, getPaymentByOrderUIDQuery, orderUID).Scan(
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

	rows, err := repo.db.QueryContext(ctx, getItemsByOrderUIDQuery, orderUID)
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

func (repo *SqlRepository) SaveOrder(ctx context.Context, order *models.Order) error {

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, createOrderQuery,
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.Shardkey,
		order.SmId,
		order.DateCreated,
		order.OofShard,
	); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.ExecContext(ctx, createDeliveryQuery,
		order.OrderUID,
		order.Name,
		order.Phone,
		order.Zip,
		order.City,
		order.Address,
		order.Region,
		order.Email,
	); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.ExecContext(ctx, createPaymentsQuery,
		order.OrderUID,
		order.Payment.Transaction,
		order.Payment.RequestId,
		order.Payment.Currency,
		order.Payment.Provider,
		order.Payment.Amount,
		order.Payment.PaymentDt,
		order.Payment.Bank,
		order.Payment.DeliveryCost,
		order.Payment.GoodsTotal,
		order.Payment.CustomFee,
	); err != nil {
		tx.Rollback()
		return err
	}

	for _, v := range order.Items {
		if _, err = tx.ExecContext(ctx, createItemQuery,
			order.OrderUID,
			v.ChrtId,
			v.TrackNumber,
			v.Price,
			v.Rid,
			v.Name,
			v.Sale,
			v.Size,
			v.TotalPrice,
			v.NmId,
			v.Brand,
			v.Status,
		); err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
