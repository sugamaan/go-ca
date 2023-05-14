package contract

// 種類をクラス化する
type freeContract struct {
}

func NewFreeContract() Contract {
	return &freeContract{}
}

// GetMaxTaskNameLength interfaceを実装する
func (f freeContract) GetMaxTaskNameLength() int {
	return 30
}

func (f freeContract) GetMaxTaskRewardAmount() uint64 {
	return 1000
}
