package value_object

import "errors"

type Reward uint64

const DefaultReward = 100

func NewReward(value uint64) (Reward, error) {
	if value < DefaultReward {
		return 0, errors.New("100より小さい報酬は設定できません")
	}
	return Reward(value), nil
}
