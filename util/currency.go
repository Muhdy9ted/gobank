package util

import (
	db "goBank/db/sqlc"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case string(db.CurrencyEUR), string(db.CurrencyUSD):
		return true
	}
	return false
}
