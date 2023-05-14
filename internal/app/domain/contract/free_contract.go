package contract

type freeContract struct {
	name         string
	price        uint32
	contractType uint32
}

func NewFreeContract(name string, price, contractType uint32) (Contract, error) {
	return &freeContract{
		name:         name,
		price:        price,
		contractType: contractType,
	}, nil
}

func (f freeContract) GetMaxTaskNameLength() int {
	return 30
}

func (f freeContract) GetMaxTaskRewardAmount() uint64 {
	return 1000
}
