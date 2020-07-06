package model

/*IGrid is the base for all the Grid implementations*/
type IGrid interface {
	Set(row int, col int, value interface{})
	Get(row int, col int) (interface{}, bool)
}
