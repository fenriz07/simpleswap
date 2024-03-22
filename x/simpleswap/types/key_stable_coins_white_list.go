package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StableCoinsWhiteListKeyPrefix is the prefix to retrieve all StableCoinsWhiteList
	StableCoinsWhiteListKeyPrefix = "StableCoinsWhiteList/value/"
)

// StableCoinsWhiteListKey returns the store key to retrieve a StableCoinsWhiteList from the index fields
func StableCoinsWhiteListKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
