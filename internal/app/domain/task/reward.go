package task

import "errors"

type Reward interface {
	Value() uint64
}

type reward struct {
	value uint64
}

const DefaultReward = 100

func NewReward(value uint64) (Reward, error) {
	if value < DefaultReward {
		return nil, errors.New("100より小さい報酬は設定できません")
	}
	return reward{value: value}, nil
}

func (r reward) Value() uint64 {
	return r.value
}
