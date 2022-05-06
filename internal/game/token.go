package game

// Token is a marker that can be put in the Table by a Player
type Token struct {
	symbol string
}

func (t Token) Symbol() string { return t.symbol }

type TokenFactory struct {
	symbol string
}

func NewTokenFactory(symbol string) *TokenFactory {
	return &TokenFactory{symbol: symbol}
}

func (f *TokenFactory) NewToken() *Token {
	return &Token{symbol: f.symbol}
}
