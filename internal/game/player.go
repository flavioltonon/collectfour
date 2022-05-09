package game

type Player struct {
	name         string
	tokenFactory *TokenFactory
}

func NewPlayer(name string, color *Color) Player {
	return Player{name: name, tokenFactory: NewTokenFactory(color)}
}

func (p *Player) Name() string { return p.name }

func (p *Player) DropToken(table *Table, column int) error {
	return table.AddToken(p.tokenFactory.NewToken(), column)
}
