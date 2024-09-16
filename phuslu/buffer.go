package phuslu

import "sync"

type byteBuffer struct {
	Bytes []byte
}

var byteBufferPool = sync.Pool{
	New: func() any {
		return &byteBuffer{
			Bytes: make([]byte, 0),
		}
	},
}

func (b *byteBuffer) Reset() {
	b.Bytes = b.Bytes[:0]
}
