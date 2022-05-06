package game

import (
	"container/ring"
	"fmt"
)

type Game struct {
	players []Player
	queue   *PlayerQueue
}

type PlayerQueue struct {
	sequence *ring.Ring
}

func NewPlayerQueue(players ...Player) *PlayerQueue {
	var queue PlayerQueue

	for i, player := range players {
		node := &ring.Ring{
			Value: player,
		}

		if i == 0 {
			queue.sequence = node
			continue
		}

		queue.sequence = queue.sequence.Link(node)
	}

	return &queue
}

func (q *PlayerQueue) NextPlayer() Player {
	current := q.sequence.Value.(Player)
	q.sequence = q.sequence.Next()
	return current
}

func New(players ...Player) *Game {
	return &Game{
		players: players,
		queue:   NewPlayerQueue(players...),
	}
}

func (g *Game) Run() error {
	for {
		var option int
		fmt.Printf("# Main menu:\n1: start match\n2: exit game\n\n")
		fmt.Printf("Select an option: ")
		fmt.Scanf("%d", &option)
		fmt.Println()
		switch option {
		case 1:
			if err := g.Start(7, 6); err != nil {
				return err
			}
		case 2:
			return nil
		}
	}
}

func (g *Game) Start(columns, rows int) error {
	table := NewTable(columns, rows)

	for table.HasAvailableColumns() {
		player := g.queue.NextPlayer()

		fmt.Printf("%s, it is your turn!\n\n", player.Name())
		table.Print()
		fmt.Println()
		for {
			var column int
			fmt.Printf("Select a column [0-%d]: ", columns-1)
			fmt.Scanf("%d", &column)
			fmt.Println()
			if err := player.DropToken(table, column); err != nil {
				fmt.Printf("Failed to drop token: %v\n", err)
			}
			break
		}

		if table.HasWinningSequence() {
			fmt.Printf("%s has own the game!\n\n", player.Name())
			table.Print()
			fmt.Println()
			break
		}
	}

	return nil
}
