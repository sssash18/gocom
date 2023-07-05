package price

import "errors"

type Price struct {
	cents    uint
	currency string
}

var(
	ErrPriceTooLow error = errors.New("Price must be greater than zero")
	ErrInvalidCurrency error = errors.New("Invalid currency")
)

func NewPrice(cents uint, currency string) (*Price, error)   {
	if cents <= 0 {
		return nil,ErrPriceTooLow
	}
	//currency codes are three chars
	if len(currency) != 3{
		return nil,ErrInvalidCurrency
	}
	return &Price{
		cents: cents,
		currency: currency,
	},nil
}

func NewPriceP(cents uint, currency string) *Price{
	price,err := NewPrice(cents,currency);
	if  err != nil{
		panic(err)
	}
	return price
}

func (p *Price) Cents() uint{
	return p.cents
}

func (p *Price) Currency() string{
	return p.currency
}