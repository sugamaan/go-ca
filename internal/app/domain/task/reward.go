package task

import (
	"errors"
	contractDomain "go-ca/internal/app/domain/contract"
)

type Reward interface {
	Value() uint64
}

type reward struct {
	value uint64
}

const (
	DefaultReward = 100
)

func NewReward(value uint64, contract contractDomain.Contract) (Reward, error) {
	if value < DefaultReward {
		return nil, errors.New("100より小さい報酬は設定できません")
	}

	// switch文を追加し契約内容ごとのバリデーションを追加する
	// まずはtypeのswitch文のアンチパターンで実装してみる。
	switch contract.ContractType() {
	case contractDomain.FreeContract:
		if value > 1000 {
			return nil, errors.New("free契約の報酬上限は1,000までです")
		}
	case contractDomain.LightContract:
		if value > 5000 {
			return nil, errors.New("light契約の報酬上限は5,000までです")
		}
	}
	return reward{value: value}, nil
}

func (r reward) Value() uint64 {
	return r.value
}
