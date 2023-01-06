package keccak

import (
	"bytes"
	"testing"
)

type testcase struct {
	msg    []byte
	output []byte
}

func TestKeccakShort256(t *testing.T) {
	h := newKeccak256()
	for i := range tstShort {
		h.Reset()
		h.Write(tstShort[i].msg)
		d := h.Sum(nil)
		if !bytes.Equal(d, tstShort[i].output) {
			t.Errorf("testcase Short256 %d: expected %x got %x", i, tstShort[i].output, d)
		}
	}
}

func TestKeccakLong256(t *testing.T) {
	h := newKeccak256()
	for i := range tstLong {
		h.Reset()
		h.Write(tstLong[i].msg)
		d := h.Sum(nil)
		if !bytes.Equal(d, tstLong[i].output) {
			t.Errorf("testcase Long256 %d: expected %x got %x", i, tstLong[i].output, d)
		}
	}
}
