package game

import (
	"fmt"
	"strings"
)

// Node is a spot for a Token in the Table
type Node struct {
	token *Token
}

func (n *Node) IsEmpty() bool { return n.token == nil }

func (n *Node) SetToken(token Token) error {
	if !n.IsEmpty() {
		return fmt.Errorf("cannot set token: node not empty")
	}

	n.token = &token
	return nil
}

func (n *Node) Color() *Color {
	if n.IsEmpty() {
		return Blank
	}

	return n.token.Color()
}

func (n Nodes) IsAvailable() bool { return n[0].IsEmpty() }

func (n Nodes) AddToken(token Token) error {
	if n.IsAvailable() {
		for i := len(n) - 1; i >= 0; i-- {
			if n[i].IsEmpty() {
				return n[i].SetToken(token)
			}
		}
	}

	return fmt.Errorf("failed to add token: no nodes available")
}

func (n Nodes) HasWinningSequence() bool {
	var (
		color *Color
		score int
	)

	for _, node := range n {
		if node.IsEmpty() {
			score = 0
			continue
		}

		if score > 0 && color == node.token.Color() {
			score++
		} else {
			color = node.token.Color()
			score = 1
		}

		if score == 4 {
			return true
		}
	}

	return false
}

const WinningScore = 4

type Nodes []*Node

type Matrix []Nodes

type Diagonals []Nodes

type Table struct {
	columns   Matrix
	rows      Matrix
	diagonals Diagonals
}

func (m Matrix) Diagonals() Diagonals {
	var (
		diagonals Diagonals
		rows      = len(m)
		columns   = len(m[0])
	)

	// Sweep direct diagonals starting from the bottom left going upwards and then to the right
	for r := rows - 1; r >= 0; r-- {
		var (
			diagonal      Nodes
			currentColumn int
		)

		for currentRow := r; currentRow < rows && currentColumn < columns; currentRow++ {
			diagonal = append(diagonal, m[currentRow][currentColumn])
			currentColumn++
		}

		diagonals = append(diagonals, diagonal)
	}

	for c := 1; c < columns; c++ {
		var (
			diagonal   Nodes
			currentRow int
		)

		for currentColumn := c; currentColumn < columns && currentRow < rows; currentColumn++ {
			diagonal = append(diagonal, m[currentRow][currentColumn])
			currentRow++
		}

		diagonals = append(diagonals, diagonal)
	}

	// Sweep secondary diagonals
	for r := 0; r < rows; r++ {
		var (
			diagonal      Nodes
			currentColumn int
		)

		for currentRow := r; currentRow >= 0 && currentColumn < columns; currentRow-- {
			diagonal = append(diagonal, m[currentRow][currentColumn])
			currentColumn++
		}

		diagonals = append(diagonals, diagonal)
	}

	for c := 1; c < columns; c++ {
		var (
			diagonal   Nodes
			currentRow = rows - 1
		)

		for currentColumn := c; currentColumn < columns && currentRow >= 0; currentColumn++ {
			diagonal = append(diagonal, m[currentRow][currentColumn])
			currentRow--
		}

		diagonals = append(diagonals, diagonal)
	}

	return diagonals
}

func NewTable(columns, rows int) *Table {
	table := &Table{
		columns:   make(Matrix, columns),
		rows:      make(Matrix, rows),
		diagonals: make([]Nodes, 0, 2*(columns+rows-1)),
	}

	for c := range table.columns {
		table.columns[c] = make(Nodes, rows)

		for r := range table.columns[c] {
			table.columns[c][r] = new(Node)
		}
	}

	for r := range table.rows {
		table.rows[r] = make(Nodes, columns)

		for c := range table.rows[r] {
			table.rows[r][c] = table.columns[c][r]
		}
	}

	table.diagonals = append(table.diagonals, table.rows.Diagonals()...)

	return table
}

func (n Nodes) Print() {
	tokens := make([]string, 0, len(n))

	for _, node := range n {
		tokens = append(tokens, node.Color().Paint("O"))
	}

	fmt.Println(strings.Join(tokens, Blank.Paint(" | ")))
}

func (t *Table) Print() {
	for _, row := range t.Rows() {
		row.Print()
	}

	fmt.Println()
}

func (t *Table) Columns() Matrix      { return t.columns }
func (t *Table) Rows() Matrix         { return t.rows }
func (t *Table) Diagonals() Diagonals { return t.diagonals }

func (t *Table) AvailableColumns() []int {
	var availableColumns []int

	for c, nodes := range t.columns {
		if nodes.IsAvailable() {
			availableColumns = append(availableColumns, c)
		}
	}

	return availableColumns
}

func (t *Table) HasAvailableColumns() bool { return len(t.AvailableColumns()) > 0 }

func (t *Table) Column(i int) (Nodes, error) {
	if i < 0 || i >= len(t.columns) {
		return nil, fmt.Errorf("column index must be between 0 and %d", len(t.columns)-1)
	}

	return t.columns[i], nil
}

func (t *Table) AddToken(token Token, c int) error {
	column, err := t.Column(c)
	if err != nil {
		return err
	}

	return column.AddToken(token)
}

func (t *Table) HasWinningSequence() bool {
	if t.columns.HasWinningSequence() {
		return true
	}

	if t.rows.HasWinningSequence() {
		return true
	}

	if t.diagonals.HasWinningSequence() {
		return true
	}

	return false
}

func (m Matrix) HasWinningSequence() bool {
	for _, nodes := range m {
		if nodes.HasWinningSequence() {
			return true
		}
	}

	return false
}

func (d Diagonals) HasWinningSequence() bool {
	for _, nodes := range d {
		if nodes.HasWinningSequence() {
			return true
		}
	}

	return false
}
