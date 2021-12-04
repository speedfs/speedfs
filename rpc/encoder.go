package rpc

import (
	"bytes"
	"encoding/binary"
)

const (
	defaultBufferSize = 4096
)

type Encoder struct {
	order binary.ByteOrder
	buf   bytes.Buffer
}

func NewEncoder() *Encoder {
	return NewEncoderOrderSize(binary.BigEndian, defaultBufferSize)
}

func NewEncoderSize(size int) *Encoder {
	return NewEncoderOrderSize(binary.BigEndian, size)
}

func NewEncoderOrderSize(order binary.ByteOrder, size int) *Encoder {
	e := &Encoder{
		order: order,
	}
	e.buf.Grow(size)
	return e
}

func (e *Encoder) Bytes() []byte {
	return e.buf.Bytes()
}

func (e *Encoder) String() string {
	return e.buf.String()
}

func (e *Encoder) Reset() {
	e.buf.Reset()
}

func (e *Encoder) EncodeBool(v bool) {
	if v {
		e.buf.WriteByte(1)
	} else {
		e.buf.WriteByte(0)
	}
}

func (e *Encoder) EncodeByte(v byte) {
	e.buf.WriteByte(v)
}

func (e *Encoder) EncodeUint8(v uint8) {
	e.buf.Write([]byte{byte(v)})
}

func (e *Encoder) EncodeUint16(v uint16) {
	var buf [2]byte
	e.order.PutUint16(buf[:], v)
	e.buf.Write(buf[:])
}

func (e *Encoder) EncodeUint32(v uint32) {
	var buf [4]byte
	e.order.PutUint32(buf[:], v)
	e.buf.Write(buf[:])
}

func (e *Encoder) EncodeUint64(v uint64) {
	var buf [8]byte
	e.order.PutUint64(buf[:], v)
	e.buf.Write(buf[:])
}

func (e *Encoder) EncodeInt8(v int8) { e.EncodeUint8(uint8(v)) }

func (e *Encoder) EncodeInt16(v int16) { e.EncodeUint16(uint16(v)) }

func (e *Encoder) EncodeInt32(v int32) { e.EncodeUint32(uint32(v)) }

func (e *Encoder) EncodeInt64(v int64) { e.EncodeUint64(uint64(v)) }

func (e *Encoder) EncodeBytes(v []byte) {
	e.buf.Write(v)
}

func (e *Encoder) EncodeString(s string) {
	e.buf.WriteString(s)
}

func (e *Encoder) EncodeMessage(msg EncoderTo) {
	msg.EncodeTo(e)
}

type EncoderTo interface {
	EncodeTo(enc *Encoder)
}
