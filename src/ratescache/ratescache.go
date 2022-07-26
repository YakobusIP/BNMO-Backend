package ratescache

type RatesCache interface {
	SetRates(key string, value float64)
	GetRates(key string) float64
}