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

	if value > contract.GetMaxTaskRewardAmount() {
		return nil, errors.New("契約プランの報酬上限を超えています")
	}
	return reward{value: value}, nil
}

func (r reward) Value() uint64 {
	return r.value
}
