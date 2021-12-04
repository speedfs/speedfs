package rpc

type Message interface {
	EncoderTo
	DecoderFrom
	Cmd() Cmd
}
