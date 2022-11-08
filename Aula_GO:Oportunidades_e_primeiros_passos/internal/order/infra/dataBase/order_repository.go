package database

import (
	"database/sql"
	"testeAula/internal/order/entity"
)

type OrderRepository struct { 
	Db *sql.DB
}
func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error { // recebe a order e salva no banco de dados
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)") // prepara a query
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)  // executa a query
	if err != nil {
		return err
	}
	return nil
}

