package model

import (
	"sync"
)

const (
	ROWS = 5000
	COLS = 3
)

/*GridModbus is the data source */
type GridModbus struct {
	data sync.Map
}

/*NewGridModbus implementation*/
func NewGridModbus() *GridModbus {
	g := new(GridModbus)
	for r := 0; r < ROWS*COLS; r++ {
		g.data.Store(r, uint16(0))
	}
	return g
}

/*Set implementation*/
func (g *GridModbus) Set(row int, col int, value interface{}) {
	switch value.(type) {
	case uint16:
		if row > ROWS || row <= 0 {
			panic("Invalid row index!")
		}
		if col > COLS || col <= 0 {
			panic("Invalid col index!")
		}
		g.data.Store((col-1)*ROWS+row-1, value)
	default:
		panic("Invalid type!")
	}
}

/*Get implementation*/
func (g *GridModbus) Get(row int, col int) (value interface{}, ok bool) {
	if row > ROWS || row <= 0 {
		panic("Invalid row index!")
	}
	if col > COLS || col <= 0 {
		panic("Invalid col index!")
	}
	value, ok = g.data.Load((col-1)*ROWS + row - 1)
	return
}
