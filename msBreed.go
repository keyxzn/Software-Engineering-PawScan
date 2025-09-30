package entities

type MsBreed struct {
	Id          uint
	Name        string
	Description string
	Origin      MsOrigin
	Size        MsSize
	Type        MsType
}
