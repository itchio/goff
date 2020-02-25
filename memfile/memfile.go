package memfile

import (
	"bytes"
	"io"
)

type FileLike interface {
	io.Writer
	io.Seeker
	io.Reader
}

func New() FileLike {
	return &memfile{}
}

type memfile struct {
	off int64
	len int64
	b   bytes.Buffer
}

func (mf *memfile) Write(p []byte) (int, error) {
	needed := int(mf.off) + len(p) - mf.b.Cap()
	if needed > 0 {
		mf.b.Grow(needed)
	}
	n := copy(mf.b.Bytes()[mf.off:], p)
	mf.off += int64(n)
	if mf.len < mf.off {
		mf.len = mf.off
	}
	return n, nil
}

func (mf *memfile) Read(p []byte) (int, error) {
	if mf.off >= mf.len {
		return 0, io.EOF
	}

	n := copy(p, mf.b.Bytes()[mf.off:mf.len])
	mf.off += int64(n)
	return n, nil
}

func (mf *memfile) Seek(off int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		mf.off = off
	case io.SeekEnd:
		mf.off = mf.len - off
	case io.SeekCurrent:
		mf.off += off
	}
	if mf.off < 0 {
		mf.off = 0
	}
	if mf.off > mf.len {
		mf.off = mf.len
	}
	return mf.off, nil
}
