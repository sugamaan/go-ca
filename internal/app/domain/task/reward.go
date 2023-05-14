package task

import "errors"

//type Reward interface {
//	Value() uint64
//}
//
//type rewardImpl struct {
//	value uint64
//}

type Reward struct {
	value uint64
}

const DefaultReward = 100

func NewReward(value uint64) (Reward, error) {
	if value < DefaultReward {
		return Reward{}, errors.New("100より小さい報酬は設定できません")
	}
	return Reward{value: value}, nil
}

//func (r rewardImpl) Value() uint64 {
//	return r.value
//}

func (r Reward) Value() uint64 {
	return r.value
}
