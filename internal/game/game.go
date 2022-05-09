package game

import (
	"container/ring"
	"fmt"
)

// PlayerQueue is a queue for the players that are playing the game
type PlayerQueue struct {
	*ring.Ring
}

// NewPlayerQueue creates a new PlayerQueue
func NewPlayerQueue(players []Player) *PlayerQueue {
	queue := PlayerQueue{
		Ring: ring.New(len(players)),
	}

	for _, player := range players {
		queue.Value = player
		queue.Ring = queue.Next()
	}

	return &queue
}

func (q *PlayerQueue) NextPlayer() Player {
	current := q.Value.(Player)
	q.Ring = q.Next()
	return current
}

func Run() {
	for {
		var option int
		fmt.Printf("# Main menu:\n1: start match\n0: exit game\n\n")
		fmt.Printf("Select an option: ")
		fmt.Scanf("%d", &option)
		fmt.Println()
		switch option {
		case 1:
			NewMatch(
				[]Player{
					NewPlayer("Player 1", Red),
					NewPlayer("Player 2", Blue),
				},
				NewTable(7, 6),
			).Start()
		case 0:
			return
		default:
			fmt.Printf("Failed to select option: %d is not a valid option\n", option)
		}
	}
}

type Match struct {
	players      []Player
	playerQueue  *PlayerQueue
	table        *Table
	columnsTotal int
}

func NewMatch(players []Player, table *Table) *Match {
	return &Match{
		players:      players,
		playerQueue:  NewPlayerQueue(players),
		table:        table,
		columnsTotal: len(table.Columns()),
	}
}

func (m *Match) Start() {
	fmt.Printf("Match started!\n\n")

	for {
		player := m.playerQueue.NextPlayer()
		fmt.Printf("%s, it is your turn!\n\n", player.Name())
		m.table.Print()

		for {
			var column int
			fmt.Printf("Select a column [1-%d]: ", m.columnsTotal)
			fmt.Scanf("%d", &column)
			fmt.Println()
			if column < 1 || column > m.columnsTotal {
				fmt.Printf("Cannot select column: column %d is out of range\n\n", column)
				continue
			}

			c, err := m.table.Column(column - 1)
			if err != nil {
				fmt.Printf("Failed to select column: %v\n\n", err)
				continue
			}

			if !c.IsAvailable() {
				fmt.Printf("Cannot select column: column %d has no spaces left\n\n", column)
				continue
			}

			if err := player.DropToken(m.table, column-1); err != nil {
				fmt.Printf("Failed to drop token: %v\n\n", err)
				continue
			}

			break
		}

		if m.table.HasWinningSequence() {
			fmt.Printf("%s has own the game!\n\n", player.Name())
			m.table.Print()
			break
		}

		if !m.table.HasAvailableColumns() {
			fmt.Printf("The game has ended with a draw!\n\n")
			m.table.Print()
			break
		}
	}
}
