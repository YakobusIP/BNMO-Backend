package symbolcache

type SymbolCache interface {
	SetSymbol(key string, value interface{})
	GetSymbol(key string) interface{}
}