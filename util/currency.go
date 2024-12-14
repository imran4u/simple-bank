package util

const (
	USD = "USD"
	RS  = "RS"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, RS:
		return true
	}
	return false
}
