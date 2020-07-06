package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGridModBus(t *testing.T) {
	gmb := NewGridModbus()

	for r := 1; r <= ROWS; r++ {
		for c := 1; c <= COLS; c++ {
			v, ok := gmb.Get(r, c)
			assert.True(t, ok, "Invalid value in GridModbus.")
			assert.Equal(t, uint16(0), v, "Invalid initialization in GridModbus.")
		}
	}
}

func TestGridModBusSetAndGet(t *testing.T) {

	var booleans [ROWS]uint16
	var floats [ROWS]uint16
	var integers [ROWS]uint16
	for i := 0; i < ROWS; i++ {
		booleans[i] = 1
		floats[i] = 123
		integers[i] = 0xFFFF
	}

	gmb := NewGridModbus()
	for c := 1; c <= COLS; c++ {
		switch c {
		case 1:
			for r := 1; r <= ROWS; r++ {
				gmb.Set(r, c, booleans[r-1])
			}
		case 2:
			for r := 1; r <= ROWS; r++ {
				gmb.Set(r, c, floats[r-1])
			}
		case 3:
			for r := 1; r <= ROWS; r++ {
				gmb.Set(r, c, integers[r-1])
			}
		default:
			panic("Invalid column.")
		}
	}

	for c := 1; c <= COLS; c++ {
		switch c {
		case 1:
			for r := 1; r <= ROWS; r++ {
				v, ok := gmb.Get(r, c)
				assert.True(t, ok, fmt.Sprintf("Invalid value in GridModbus row: %v, col: %v", r, c))
				assert.Equal(t, uint16(1), v, fmt.Sprintf("Invalid value in GridModbus row: %v, col: %v", r, c))
			}
		case 2:
			for r := 1; r <= ROWS; r++ {
				v, ok := gmb.Get(r, c)
				assert.True(t, ok, fmt.Sprintf("Invalid value in GridModbus row: %v, col: %v", r, c))
				assert.Equal(t, uint16(123), v, fmt.Sprintf("Invalid value in GridModbus row: %v, col: %v", r, c))
			}
		case 3:
			for r := 1; r <= ROWS; r++ {
				v, ok := gmb.Get(r, c)
				assert.True(t, ok, fmt.Sprintf("Invalid value in GridModbus row: %v, col: %v", r, c))
				assert.Equal(t, uint16(0xFFFF), v, fmt.Sprintf("Invalid value in GridModbus row: %v, col: %v", r, c))
			}
		default:
			panic("Invalid column.")
		}
	}
}

func TestGridModBusInvalidGetRow(t *testing.T) {
	gmb := NewGridModbus()
	defer func() { recover() }()
	// This function should cause a panic
	gmb.Get(0, 1)
	t.Errorf("gmb.Get(0, 1) should have panicked!")
}

func TestGridModBusInvalidGetCol(t *testing.T) {
	gmb := NewGridModbus()
	defer func() { recover() }()
	// This function should cause a panic
	gmb.Get(1, 0)
	t.Errorf("gmb.Get(1, 0) should have panicked!")
}

func TestGridModBusInvalidSetRow(t *testing.T) {
	gmb := NewGridModbus()
	defer func() { recover() }()
	// This function should cause a panic
	gmb.Set(0, 1, uint16(1))
	t.Errorf("gmb.Set(0, 1) should have panicked!")
}

func TestGridModBusInvalidSetCol(t *testing.T) {
	gmb := NewGridModbus()
	defer func() { recover() }()
	// This function should cause a panic
	gmb.Set(1, 0, uint16(1))
	t.Errorf("gmb.Set(1, 0) should have panicked!")
}
func TestGridModBusInvalidSetValue(t *testing.T) {
	gmb := NewGridModbus()
	defer func() { recover() }()
	// This function should cause a panic
	gmb.Set(1, 1, int64(1))
	t.Errorf("gmb.Set(1, 0) should have panicked!")
}
