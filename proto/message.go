package proto

type Message interface {
	EncoderTo
	DecoderFrom
}
