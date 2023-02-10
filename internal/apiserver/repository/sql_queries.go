package repository

const (
	getOrderByOrderUIDQuery = `SELECT order_uid, track_number, entry, locale,
       internal_signature, customer_id,delivery_service,shardkey, sm_id,
       date_created, oof_shard FROM orders WHERE order_uid = $1`
	getItemsByOrderUIDQuery = `SELECT chrt_id, track_number, price, rid, name,
       sale, size, total_price, nm_id, brand, status FROM items WHERE order_uid = $1`
	getPaymentByOrderUIDQuery = `SELECT transaction, request_id, currency, provider,
    amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payments WHERE order_uid = $1`
	getDeliveryByOrderUIDQuery = `SELECT name, phone, zip, city, address,
       region, email FROM deliveries where order_uid = $1`
)
