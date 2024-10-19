package common_test

import (
	"testing"

	"github.com/yairp7/go-common-lib/common"
	"gopkg.in/stretchr/testify.v1/assert"
)

const (
	SAMPLE_16bit int64 = 15543           // 0x3CB7
	SAMPLE_32bit int64 = 15543522        // 0xED2CE2
	SAMPLE_64bit int64 = 155438585334556 // 0x8D5EDDA9CB1C
)

func Test_Int2Bytes_and_Bytes2Int(t *testing.T) {
	bytes := common.Int2Bytes(SAMPLE_16bit, 16)
	assert.Equal(t, byte(0xB7), bytes[0])
	assert.Equal(t, byte(0x3C), bytes[1])

	v := common.Bytes2Int(bytes, 16)
	assert.Equal(t, SAMPLE_16bit, v)

	bytes = common.Int2Bytes(SAMPLE_32bit, 32)
	assert.Equal(t, byte(0xE2), bytes[0])
	assert.Equal(t, byte(0x2C), bytes[1])
	assert.Equal(t, byte(0xED), bytes[2])

	v = common.Bytes2Int(bytes, 32)
	assert.Equal(t, SAMPLE_32bit, v)

	bytes = common.Int2Bytes(SAMPLE_64bit, 64)
	assert.Equal(t, byte(0x1C), bytes[0])
	assert.Equal(t, byte(0xCB), bytes[1])
	assert.Equal(t, byte(0xA9), bytes[2])
	assert.Equal(t, byte(0xDD), bytes[3])
	assert.Equal(t, byte(0x5E), bytes[4])
	assert.Equal(t, byte(0x8D), bytes[5])

	v = common.Bytes2Int(bytes, 64)
	assert.Equal(t, SAMPLE_64bit, v)
}
