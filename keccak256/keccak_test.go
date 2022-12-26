package keccak

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestKeccak256(t *testing.T) {
	h := newKeccak256()
	h.Reset()
	h.Write([]byte("The quick brown fox jumps over the lazy dog, The quick brown fox jumps over the lazy dog, The quick brown fox jumps over the lazy dog, The quick brown fox jumps over the lazy dog, The quick brown fox jumps over the lazy dog, The quick brown fox jumps over the lazy dog"))
	d := h.Sum(nil)
	expected := "e4b5f1d230b0f84b4c7261a7b078cef6d3a7af18256f58d6f817658fd7bc08a2"
	if !bytes.Equal(d, common.Hex2Bytes(expected)) {
		t.Errorf("Testcase Keccak256 failed, expected %x got %x", expected, d)
	}
}
