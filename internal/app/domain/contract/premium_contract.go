package contract

type premiumContract struct {
	name         string
	price        uint32
	contractType uint32
}

func NewPremiumContract(name string, price uint32, contractType uint32) (Contract, error) {
	return &premiumContract{
		name:         name,
		price:        price,
		contractType: contractType,
	}, nil
}
