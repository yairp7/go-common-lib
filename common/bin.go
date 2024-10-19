package common

import "encoding/binary"

func Int2Bytes(n int64, bits int64) []byte {
	bs := make([]byte, bits/8)
	switch bits {
	case 16:
		binary.LittleEndian.PutUint16(bs, uint16(n))
		break
	case 32:
		binary.LittleEndian.PutUint32(bs, uint32(n))
		break
	case 64:
		binary.LittleEndian.PutUint64(bs, uint64(n))
		break
	}
	return bs
}

func Bytes2Int(b []byte, bits int) int64 {
	switch bits {
	case 16:
		return int64(binary.LittleEndian.Uint16(b))
	case 32:
		return int64(binary.LittleEndian.Uint32(b))
	case 64:
		return int64(binary.LittleEndian.Uint64(b))
	}
	return 0
}
