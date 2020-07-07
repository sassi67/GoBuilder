package model

import "log"

/*BoardEVO implements IBoard*/
type BoardEVO struct {
	grid    *GridModbus
	mapData map[int]GridMap
}

/*NewBoardEVO : BoardEVO constructor*/
func NewBoardEVO() (*BoardEVO, int, error) {
	newGridModbus := NewGridModbus()
	newBoardEVO := &BoardEVO{grid: newGridModbus}

	size, err := newBoardEVO.initMap()
	if err != nil {
		log.Fatal("initMap error")
		return nil, 0, err
	}

	return newBoardEVO, size, nil
}

func (evo *BoardEVO) initMap() (int, error) {
	// data association
	evo.mapData = map[int]GridMap{
		BoardType:  GridMap{Row: 1, Col: 3, Size: 1},
		Season:     GridMap{Row: 1, Col: 1, Size: 1},
		BoardOnOff: GridMap{Row: 2, Col: 1, Size: 1},
		TempExt:    GridMap{Row: 1, Col: 2, Size: 1},
	}

	// set the specific value for BoardType
	evo.grid.Set(1, 3, uint16(BoardEvo))

	return len(evo.mapData), nil
}

/*GetBoardType returns the board family*/
func (evo *BoardEVO) GetBoardType() (uint16, bool) {
	v, ok := evo.grid.Get(1, 3)

	switch v.(type) {
	case uint16:
		return v.(uint16), ok
	default:
	}

	return 0, false
}
