package contract

type lightContract struct {
	name         string
	price        uint32
	contractType uint32
}

func newLightContract(name string, price uint32, contractType uint32) (Contract, error) {
	return &lightContract{
		name:         name,
		price:        price,
		contractType: contractType,
	}, nil
}

func (l lightContract) GetMaxTaskNameLength() int {
	return 50
}

func (l lightContract) GetMaxTaskRewardAmount() uint64 {
	return 5000
}
