package contract

type lightContract struct {
}

func NewLightContract() Contract {
	return &lightContract{}
}

func (l lightContract) GetMaxTaskNameLength() int {
	return 50
}

func (l lightContract) GetMaxTaskRewardAmount() uint64 {
	return 5000
}
