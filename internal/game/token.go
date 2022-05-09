package game

import "github.com/fatih/color"

// Token is a marker that can be put in the Table by a Player
type Token struct {
	color *Color
}

func (t *Token) Color() *Color { return t.color }

type TokenFactory struct {
	color *Color
}

type Color struct {
	painter *color.Color
}

func (c *Color) Paint(format string, a ...interface{}) string {
	return c.painter.Sprintf(format, a...)
}

var (
	Blank = &Color{painter: color.New(color.FgWhite)}
	Red   = &Color{painter: color.New(color.FgRed)}
	Blue  = &Color{painter: color.New(color.FgBlue)}
)

func NewTokenFactory(color *Color) *TokenFactory {
	return &TokenFactory{color: color}
}

func (f *TokenFactory) NewToken() Token {
	return Token{color: f.color}
}
