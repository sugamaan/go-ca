package contract

import "errors"

const (
	FreeContractType   = 1
	FreeContractName   = "無料プラン"
	FreeContractPrice  = 0
	LightContractType  = 2
	LightContractName  = "ライトプラン"
	LightContractPrice = 1000
)

type Contract interface {
	GetMaxTaskNameLength() int
	GetMaxTaskRewardAmount() uint64
}

func NewContract(contractType uint32) (Contract, error) {
	if contractType != FreeContractType && contractType != LightContractType {
		return nil, errors.New("契約タイプが不正です")
	}
	switch contractType {
	case FreeContractType:
		return newFreeContract(FreeContractName, FreeContractPrice, contractType)
	case LightContractType:
		return newLightContract(LightContractName, LightContractPrice, contractType)
	}
	return nil, errors.New("契約タイプが対象外です")
}
