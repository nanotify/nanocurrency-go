package nanocurrency

import (
	"golang.org/x/crypto/blake2b"
)

const checksumLength = 5

func reverseBytes(xs []byte) {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
}

func computeChecksum(xs []byte) ([]byte, error) {
	hasher, err := blake2b.New(checksumLength, nil)
	if err != nil {
		return nil, err
	}

	_, err = hasher.Write(xs)
	if err != nil {
		return nil, err
	}

	res := hasher.Sum(nil)

	reverseBytes(res)

	return res, nil
}
