package proto

type Header struct {
	BodyLen uint8
	Cmd     uint8
	Status  int8
}
