package contract

import "errors"

const (
	FreeContractType    = 1
	LightContractType   = 2
	PremiumContractType = 3
)

// Contract 切り替えたい機能をinterfaceのメソッドとして定義する
type Contract interface {
	GetMaxTaskNameLength() int
	GetMaxTaskRewardAmount() uint64
}

type contract struct {
	name         string
	price        uint32
	contractType uint32
}

func NewContract(name string, price uint32, contractType uint32) (Contract, error) {
	if contractType != FreeContractType && contractType != LightContractType && contractType != PremiumContractType {
		return nil, errors.New("契約タイプが不正です")
	}
	switch contractType {
	case FreeContractType:
		return NewFreeContract(), nil
	case LightContractType:
		return NewLightContract(), nil
		//case PremiumContractType:
		//	return &premiumContract{}, nil
	}
	return nil, errors.New("契約タイプが対象外です")
}

//
//func (c contract) Name() string {
//	return c.name
//}
//
//func (c contract) Price() uint32 {
//	return c.price
//}
//
//func (c contract) ContractType() uint32 {
//	return c.contractType
//}
