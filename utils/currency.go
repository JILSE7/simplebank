package utils

// Constansts for all supported curerncies
const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
)

func IsValidCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}

	return false
}
