package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type HashWriter struct {
	hashObj hash.Hash
}

func (h HashWriter) Write(b []byte) (n int, err error) {
	_, err = h.hashObj.Write(b)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (h HashWriter) Result() []byte {
	return h.hashObj.Sum(nil)
}

func (h HashWriter) ResultAsHex() []byte {
	return []byte(hex.EncodeToString(h.hashObj.Sum(nil)))
}

func SHA256Writer() HashWriter {
	return HashWriter{hashObj: sha256.New()}
}

func SHA1Writer() HashWriter {
	return HashWriter{hashObj: sha1.New()}
}

func MD5Writer() HashWriter {
	return HashWriter{hashObj: md5.New()}
}

func hashBytes(hashWriter HashWriter, data []byte, asHex bool) ([]byte, error) {
	_, err := hashWriter.Write(data)
	if err != nil {
		return nil, err
	}
	if asHex {
		return hashWriter.ResultAsHex(), nil
	}
	return hashWriter.Result(), nil
}

func SHA256(data []byte, asHex bool) ([]byte, error) {
	return hashBytes(SHA256Writer(), data, asHex)
}

func SHA1(data []byte, asHex bool) ([]byte, error) {
	return hashBytes(SHA1Writer(), data, asHex)
}

func MD5(data []byte, asHex bool) ([]byte, error) {
	return hashBytes(MD5Writer(), data, asHex)
}
