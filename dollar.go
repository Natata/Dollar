package dollar

import "github.com/shopspring/decimal"

type Dollar struct {
	decimal.Decimal
}

func (d *Dollar) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Dollar) UnmarshalJSON(b []byte) error {
	amount, err := decimal.NewFromString(string(b))
	if err != nil {
		return nil
	}
	d.Decimal = amount
	return nil
}
