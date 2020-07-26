package ioutils

import (
	"bytes"
	"io"
)

// LoopReader type define
type LoopReader struct {
	eof  bool
	data []byte
	*bytes.Reader
}

// NewLoopReader returns a new LoopReader reading from data.
func NewLoopReader(data []byte) *LoopReader {
	return NewLoopReaderEOF(data, true)
}

// NewLoopReaderEOF returns a new LoopReader from data and eof.
func NewLoopReaderEOF(data []byte, eof bool) *LoopReader {
	return &LoopReader{
		eof:    eof,
		data:   data,
		Reader: bytes.NewReader(data),
	}
}

// Read implements the io.Reader interface.
func (r *LoopReader) Read(b []byte) (n int, err error) {
	if n, err = r.Reader.Read(b); err == io.EOF {
		if r.Reader.Reset(r.data); !r.eof {
			err = nil
		}
	}
	return
}

// ReadAt implements the io.Reader interface.
func (r *LoopReader) ReadAt(b []byte, off int64) (n int, err error) {
	if n, err = r.Reader.ReadAt(b, off); err == io.EOF {
		if r.Reader.Reset(r.data); !r.eof {
			err = nil
		}
	}
	return
}

// ReadByte implements the io.Reader interface.
func (r *LoopReader) ReadByte() (b byte, err error) {
	if b, err = r.Reader.ReadByte(); err == io.EOF {
		if r.Reader.Reset(r.data); !r.eof {
			err = nil
		}
	}
	return
}

// ReadRune implements the io.Reader interface.
func (r *LoopReader) ReadRune() (ch rune, size int, err error) {
	if ch, size, err = r.Reader.ReadRune(); err == io.EOF {
		if r.Reader.Reset(r.data); !r.eof {
			err = nil
		}
	}
	return
}
