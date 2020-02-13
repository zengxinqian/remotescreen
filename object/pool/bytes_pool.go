package pool

import "sync"

var bytesPool sync.Pool

func GetBytes(length int) []byte {

	buffer := bytesPool.Get().([]byte)
	if buffer == nil {
		return make([]byte, length)
	}
	if cap(buffer) < length {
		bytesPool.Put(buffer)
		return make([]byte, length)
	}
	return buffer

}

func PutBytes(bytes []byte) {
	bytesPool.Put(bytes)
}
