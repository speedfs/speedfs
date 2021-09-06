package rpc

type Message interface {
	EncoderTo
	DecoderFrom
}
