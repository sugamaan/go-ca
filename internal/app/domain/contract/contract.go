package contract

import "errors"

const (
	FreeContractType  = 1
	LightContractType = 2
)

type Contract interface {
	GetMaxTaskNameLength() int
	GetMaxTaskRewardAmount() uint64
}

func NewContract(name string, price uint32, contractType uint32) (Contract, error) {
	if contractType != FreeContractType && contractType != LightContractType {
		return nil, errors.New("契約タイプが不正です")
	}
	switch contractType {
	case FreeContractType:
		return NewFreeContract(name, price, contractType)
	case LightContractType:
		return NewLightContract(name, price, contractType)
	}
	return nil, errors.New("契約タイプが対象外です")
}
