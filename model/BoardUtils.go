package model

/* Board types enumeration: CRC, EVO and WI */
const (
	BoardCrc = iota
	BoardEvo
	BoardWI
)

/*GridMap is a tuple which is the value */
type GridMap struct {
	Row  int
	Col  int
	Size int
}

/* Data type is the key a GridMap*/
const (
	BoardType = iota
	Season
	BoardOnOff
	TempExt
	DataEnd /* keep it as last element in this list! */
)
