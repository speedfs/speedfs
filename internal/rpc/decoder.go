package rpc

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Decoder struct {
	buf   *bytes.Buffer
	order binary.ByteOrder
}

func NewDecoder(buf []byte) *Decoder {
	return NewDecoderOrder(buf, binary.BigEndian)
}

func NewDecoderOrder(buf []byte, order binary.ByteOrder) *Decoder {
	d := &Decoder{
		buf:   bytes.NewBuffer(buf),
		order: order,
	}
	return d
}

func (d *Decoder) Reset() {
	d.buf.Reset()
}

func (d *Decoder) DecodeBytes(v []byte) error {
	n, err := d.buf.Read(v)
	if err != nil {
		return err
	}
	if n < len(v) {
		return io.EOF
	}

	return nil
}

func (d *Decoder) DecodeBool() (bool, error) {
	u, err := d.DecodeUint8()
	if err != nil {
		return false, err
	}
	return u != 0, nil
}

func (d *Decoder) DecodeUint8() (uint8, error) {
	var buf [1]byte
	if err := d.DecodeBytes(buf[:]); err != nil {
		return 0, err
	}
	return buf[0], nil
}

func (d *Decoder) DecodeUint16() (uint16, error) {
	var buf [2]byte
	if err := d.DecodeBytes(buf[:]); err != nil {
		return 0, err
	}
	return d.order.Uint16(buf[:]), nil
}

func (d *Decoder) DecodeUint32() (uint32, error) {
	var buf [4]byte
	if err := d.DecodeBytes(buf[:]); err != nil {
		return 0, err
	}
	return d.order.Uint32(buf[:]), nil
}

func (d *Decoder) DecodeUint64() (uint64, error) {
	var buf [8]byte
	if err := d.DecodeBytes(buf[:]); err != nil {
		return 0, err
	}
	return d.order.Uint64(buf[:]), nil
}

func (d *Decoder) DecodeInt8() (int8, error) {
	u, err := d.DecodeUint8()
	return int8(u), err
}

func (d *Decoder) DecodeInt16() (int16, error) {
	u, err := d.DecodeUint16()
	return int16(u), err
}

func (d *Decoder) DecodeInt32() (int32, error) {
	u, err := d.DecodeUint32()
	return int32(u), err
}

func (d *Decoder) DecodeInt64() (int64, error) {
	u, err := d.DecodeUint64()
	return int64(u), err
}

func (d *Decoder) DecodeMessage(msg DecoderFrom) error {
	return msg.DecodeFrom(d)
}

type DecoderFrom interface {
	DecodeFrom(dec *Decoder) error
}
