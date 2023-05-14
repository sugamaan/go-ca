package contract

type freeContract struct {
}

func NewFreeContract() Contract {
	return &freeContract{}
}

func (f freeContract) GetMaxTaskNameLength() int {
	return 30
}

func (f freeContract) GetMaxTaskRewardAmount() uint64 {
	return 1000
}
