package contract

import "errors"

const (
	FreeContract    = 1
	LightContract   = 2
	PremiumContract = 3
)

type Contract interface {
	Name() string
	Price() uint32
	ContractType() uint32
}

type contract struct {
	name         string
	price        uint32
	contractType uint32
}

func NewContract(name string, price uint32, contractType uint32) (Contract, error) {
	if contractType != FreeContract && contractType != LightContract && contractType != PremiumContract {
		return nil, errors.New("契約タイプが不正です")
	}
	return contract{name: name, price: price, contractType: contractType}, nil
}

func (c contract) Name() string {
	return c.name
}

func (c contract) Price() uint32 {
	return c.price
}

func (c contract) ContractType() uint32 {
	return c.contractType
}
