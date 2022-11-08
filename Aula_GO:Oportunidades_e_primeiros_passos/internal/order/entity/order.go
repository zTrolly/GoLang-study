package entity

import "errors"

type Order struct { // tipo de uma struct
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) { // quando criar retorna um order e um error
	order := &Order{
		ID:         id,
		Price:      price,
		Tax:        tax,
	}
	err := order.IsValid() // valida se o order é valido
		if err != nil {
			return nil, err
		}
	
	return order, nil
}

func (o *Order) IsValid() error { // verifica se o order é valido olhando todos os campos
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
		
	}

	return nil
} 

func (o *Order) CalculatePrice()  error{ // calcula o preço final
	o.FinalPrice = o.Price + o.Tax
	err :=  o.IsValid()
	if err != nil {
		return err
	}
	return nil
}