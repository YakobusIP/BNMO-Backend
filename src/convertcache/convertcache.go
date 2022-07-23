package convertcache

type ConvertCache interface {
	SetConvert(key string, value int)
	GetConvert(key string) int
}