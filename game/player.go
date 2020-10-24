package game

// Player é do Tipo do jogador que realizará as jogadas
type Player struct {
	Name   string
	Symbol string
	Number int
	Winner bool
}

// NewPlayer cria novo jogador
func NewPlayer(name string, symbol string) *Player {
	p := Player{Name: name, Symbol: symbol}
	if symbol == "X" {
		p.Number = 1
	} else {
		p.Number = 10
	}
	p.Winner = false
	return &p
}
