package model

/*IBoard is the base for all the Board implementations*/
type IBoard interface {
	initMap() (int, error)
	Get(datatype uint16) (uint16, error)
	Set(datatype uint16, value uint16) error
	GetBoardType() (uint16, bool)
}
