package sha256

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func TestSha2256(t *testing.T) {
	msg := []byte("The quick brown fox jumps over the lazy dog, The quick brown fox jumps over the lazy dog, The quick brown fox jumps over the lazy dog, The quick brown fox jumps over the lazy dog, The quick brown fox jumps over the lazy dog, The quick brown fox jumps over the lazy dog")
	h := Sha256(msg)
	hash := sha256.New()
	hash.Reset()
	hash.Write(msg)
	d := hash.Sum(nil)
	if !bytes.Equal(h, d) {
		t.Errorf("Testcase Keccak256 failed, expected %x got %x", d, h)
	}
}
