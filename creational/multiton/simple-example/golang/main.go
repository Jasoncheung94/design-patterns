package main

import (
	"bytes"
	"fmt"
	"sync"
)

var pool = sync.Pool{
	// Creates an instance of a buffer if none are available to return.
	// Must return an interface{} to make it flexible and cast your type when retrieving.
	New: func() interface{} {
		// Return an instance of bytes.Buffer
		return &bytes.Buffer{}
	},
}

func main() {
	// When getting from a Pool, you need to cast
	s := pool.Get().(*bytes.Buffer)
	s.Write([]byte("Text written to buffer"))
	// Put instance back into pool for future use.
	pool.Put(s)

	// Pools can return dirty results, so you should always reset
	// Get another pool buffer.
	s = pool.Get().(*bytes.Buffer)
	s.Write([]byte(":Adding more text without a reset ? Dirty buffer."))
	// If GC ran, then the buffer would be cleaned up otherwise it would contain existing data.
	fmt.Println(s)
	//Clean up buffer before putting back into pool.
	s.Reset()
	pool.Put(s)

	// When you clean up, your buffer should be empty
	s = pool.Get().(*bytes.Buffer)
	defer pool.Put(s)
	s.Write([]byte("reset!"))
	fmt.Println(s)
	s.Reset()
}
