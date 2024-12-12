package lock_free

import (
	"bytes"
	"fmt"
	"sync"
)

type printId interface {
	print(id int)
}
type BufferWrapper struct {
	buffer *bytes.Buffer
}

func NewBufferWrapper() *BufferWrapper {
	return &BufferWrapper{}
}

// 定义一个sync.Pool，用于管理缓冲区对象
var bufferPool = sync.Pool{
	New: func() interface{} {
		// 初始化一个新的bytes.Buffer
		return &BufferWrapper{
			buffer: new(bytes.Buffer),
		}
	},
}

func (b *BufferWrapper) getBuffer() *BufferWrapper {
	return bufferPool.Get().(*BufferWrapper)
}
func (b *BufferWrapper) print(id int) {
	buf := b.getBuffer()
	defer b.putBuffer(buf)
	buf.buffer.WriteString(fmt.Sprintf("Goroutine %d\n", id))
}

// reset buffer
func (b *BufferWrapper) putBuffer(buf *BufferWrapper) {
	buf.buffer.Reset()
	bufferPool.Put(buf)
}

type SingleBuffer struct {
	buffer *bytes.Buffer
	lock   sync.Mutex
}

var singleBuff *SingleBuffer

func init() {
	singleBuff = &SingleBuffer{
		buffer: new(bytes.Buffer),
		lock:   sync.Mutex{},
	}
}

func (s *SingleBuffer) print(id int) {
	s.lock.Lock()
	s.buffer.WriteString(fmt.Sprintf("Goroutine %d\n", id))
	s.buffer.Reset()
	s.lock.Unlock()
}
