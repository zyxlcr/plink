package iface

import (
	"math/rand"
	"sync"
	"time"
)

var (
	idMutex     sync.Mutex
	lastID      uint32
	maxSequence uint16 = 0xFFFF
)

func GenerateMessageID() uint32 {
	idMutex.Lock()
	defer idMutex.Unlock()

	now := uint32(time.Now().UnixNano() / int64(time.Millisecond))

	if now == lastID {
		sequence := uint16(rand.Intn(int(maxSequence)))
		return (now << 16) | uint32(sequence)
	}

	lastID = now
	return (now << 16)
}
