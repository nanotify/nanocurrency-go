package nanocurrency_test

import (
	"testing"

	"github.com/nanotify/nanocurrency-go/nanocurrency"
	"github.com/stretchr/testify/assert"
)

func TestIsValidAccount(t *testing.T) {
	testCases := []struct {
		Description string
		Address     string
		Valid       bool
	}{
		{
			Description: "valid xrb_ address",
			Address: "xrb_" +
				"1xj5wdjge3isw97waep7fygr71q3n8ybxsrae6n5whdh31et3pdiqosi1xsk",
			Valid: true,
		},
		{
			Description: "valid nano_ address",
			Address: "nano_" +
				"1xj5wdjge3isw97waep7fygr71q3n8ybxsrae6n5whdh31et3pdiqosi1xsk",
			Valid: true,
		},
		{
			Description: "invalid foo_ address",
			Address: "foo_" +
				"1xj5wdjge3isw97waep7fygr71q3n8ybxsrae6n5whdh31et3pdiqosi1xsk",
			Valid: false,
		},
		{
			Description: "invalid checksum",
			Address: "nano_" +
				"1xj5wdjge3isw97waep7fygr71q3n8ybxsrae6n5whdh31et3pdiqosi1ssx",
			Valid: false,
		},
	}

	asserts := assert.New(t)

	for _, test := range testCases {
		res := nanocurrency.IsValidAccount(test.Address)
		asserts.Equal(test.Valid, res, test.Description)
	}
}
