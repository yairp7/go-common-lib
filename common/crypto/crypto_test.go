package crypto_test

import (
	"testing"

	"github.com/yairp7/go-common-lib/common/crypto"
	"gopkg.in/stretchr/testify.v1/assert"
)

func Test_Crypto_SHA1(t *testing.T) {
	input := "crocodile"
	expected := "c1673dcace5268cf115e28db8e84920dd5bb38ab"
	actual, err := crypto.SHA1([]byte(input), true)
	assert.Nil(t, err)
	assert.Equal(t, expected, string(actual))
}
