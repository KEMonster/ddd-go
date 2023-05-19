package entity

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Order struct {
	OrderId  uuid.UUID
	User     User
	Products []Product
}

func (o *Order) ToPO() OrderPo {
	products, _ := json.Marshal(o.Products)
	addres, _ := json.Marshal(o.User.Addr)
	return OrderPo{
		Id:         o.OrderId,
		UserId:     o.User.Uid,
		ProductIds: string(products),
		Address:    string(addres),
	}
}
