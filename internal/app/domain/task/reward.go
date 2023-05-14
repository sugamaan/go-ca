package task

import "errors"

type Reward interface {
	Value() uint64
}

type rewardImpl struct {
	value uint64
}

const DefaultReward = 100

func NewReward(value uint64) (Reward, error) {
	if value < DefaultReward {
		return nil, errors.New("100より小さい報酬は設定できません")
	}
	return rewardImpl{value: value}, nil
}

func (r rewardImpl) Value() uint64 {
	return r.value
}
