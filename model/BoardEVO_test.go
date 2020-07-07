package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoardEVO(t *testing.T) {
	assertion := assert.New(t)

	evo, size, err := NewBoardEVO()
	assertion.NotNil(evo, "No new BoardEVO created")
	assertion.Nil(err, "Error in NewBoardEVO")
	assertion.Equal(DataEnd, size, "Error in data initialization")

	family, ok := evo.GetBoardType()
	assertion.True(ok, "Error in family detection")
	assertion.Equal(uint16(BoardEvo), family, "Wrong board family")
}
