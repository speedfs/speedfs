package proto

type Header struct {
	BodyLen uint64
	Cmd     uint8
	Status  int8
}
