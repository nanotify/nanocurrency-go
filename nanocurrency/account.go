// Package nanocurrency contains the main toolkit for the Nano cryptocurrency
package nanocurrency

import (
	"encoding/base32"
	"fmt"
)

// Account encapsulates a Nano account, including it's address and other
// details.
type Account struct {
	Address string
}

const (
	nanoPrefixLength = 65
	xrbPrefixLength  = 64

	nanoPrefixStart = 5
	xrbPrefixStart  = 4

	publicKeyTop = 52
)

var (
	// ErrorInvalidAccountPrefix describes an invalid prefix for the
	// Nano account. This means the account did not start with either
	// xrb_ or nano_
	ErrorInvalidAccountPrefix = fmt.Errorf("invalid account prefix")

	// ErrorChecksumMismatch describes an error in with the checksum of an
	// address does not match the computed checksum for the public key bytes.
	ErrorChecksumMismatch = fmt.Errorf("checksum mismatch")

	// ErrorInvalidLength describes an error in which the length of the
	// account string is invalid.
	ErrorInvalidLength = fmt.Errorf("invalid account length")
)

// NewAccount attempts to initialize an Account from the string
// input paramater. If the string is invalid, then an error is returned.
func NewAccount(address string) (*Account, error) {
	switch len(address) {
	case nanoPrefixLength:
		if address[:nanoPrefixStart] != "nano_" {
			return nil, ErrorInvalidAccountPrefix
		}

		address = address[nanoPrefixStart:]
	case xrbPrefixLength:
		if address[:xrbPrefixStart] != "xrb_" {
			return nil, ErrorInvalidAccountPrefix
		}

		address = address[xrbPrefixStart:]
	default:
		return nil, ErrorInvalidLength
	}

	encoder := base32.NewEncoding("13456789abcdefghijkmnopqrstuwxyz")

	publicKeyBytes, err := encoder.DecodeString("1111" + address[:publicKeyTop])
	if err != nil {
		return nil, err
	}

	checksum, err := computeChecksum(publicKeyBytes[3:])
	if err != nil {
		return nil, err
	}

	if encoder.EncodeToString(checksum) != address[publicKeyTop:] {
		return nil, ErrorChecksumMismatch
	}

	return &Account{
		Address: address,
	}, nil
}
