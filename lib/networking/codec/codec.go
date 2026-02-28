package codec

import (
	"bytes"
	"encoding/binary"
	"io"
	"strings"
)

type PacketReader struct {
	r io.Reader
}

type PacketWriter struct {
	w io.Writer
}

func NewPacketReader(r io.Reader) *PacketReader {
	if r == nil {
		r = bytes.NewBuffer(nil)
	}
	return &PacketReader{r: r}
}

func NewPacketWriter(w io.Writer) *PacketWriter {
	if w == nil {
		w = bytes.NewBuffer(nil)
	}
	return &PacketWriter{w: w}
}

func (w *PacketWriter) Byte(v uint8) error {
	return binary.Write(w.w, binary.BigEndian, v)
}

func (r *PacketReader) Byte() (uint8, error) {
	var v uint8
	err := binary.Read(r.r, binary.BigEndian, &v)
	return v, err
}

func (w *PacketWriter) FByte(v float32) error {
	raw := int8(v * 32)
	return binary.Write(w.w, binary.BigEndian, raw)
}

func (r *PacketReader) FByte() (float32, error) {
	var raw int8
	err := binary.Read(r.r, binary.BigEndian, &raw)
	return float32(raw) / 32.0, err
}

func (w *PacketWriter) String64(s string) error {
	buf := make([]byte, 64)
	copy(buf, s)
	for i := len(s); i < 64; i++ {
		buf[i] = 0x20
	}
	_, err := w.w.Write(buf)
	return err
}

func (r *PacketReader) String64() (string, error) {
	buf := make([]byte, 64)
	_, err := io.ReadFull(r.r, buf)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(buf), " "), nil
}

func (w *PacketWriter) Short(v int16) error {
	return binary.Write(w.w, binary.BigEndian, v)
}

func (r *PacketReader) Short() (int16, error) {
	var v int16
	err := binary.Read(r.r, binary.BigEndian, &v)
	return v, err
}

func (w *PacketWriter) SByte(v int8) error {
	return binary.Write(w.w, binary.BigEndian, v)
}

func (r *PacketReader) SByte() (int8, error) {
	var v int8
	err := binary.Read(r.r, binary.BigEndian, &v)
	return v, err
}

func (w *PacketWriter) FShort(v float32) error {
	raw := int16(v * 32)
	return binary.Write(w.w, binary.BigEndian, raw)
}

func (r *PacketReader) FShort() (float32, error) {
	var raw int16
	err := binary.Read(r.r, binary.BigEndian, &raw)
	return float32(raw) / 32.0, err
}

func (w *PacketWriter) Bytes(data []byte) error {
	_, err := w.w.Write(data)
	return err
}

func (r *PacketReader) Bytes(length int) ([]byte, error) {
	buf := make([]byte, length)
	_, err := io.ReadFull(r.r, buf)
	return buf, err
}
