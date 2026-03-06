package money

import "errors"

// Money représente un montant monétaire en centimes (int64).
type Money int64

// ErrInvalidAmount est retourné si le montant est ≤ 0.
var ErrInvalidAmount = errors.New("amount must be greater than zero")

// NewMoney crée un montant validé.
func NewMoney(cents int64) (Money, error) {
	if cents <= 0 {
		return 0, ErrInvalidAmount
	}
	return Money(cents), nil
}

// Int64 retourne la valeur brute en centimes.
func (m Money) Int64() int64 {
	return int64(m)
}

// String retourne une représentation lisible (ex: "1050 cents").
func (m Money) String() string {
	return fmt.Sprintf("%d cents", m)
}