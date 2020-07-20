package noun

import (
	"hash/crc32"
	"sync"
)

type segMap struct {
	items []*bucket
}

type bucket struct {
	sync.RWMutex
	bItems map[string]*item
}

func (s *segMap) Put(item *item) {
	tar := s.items[s.hashKeyM(item.key)]
	tar.Lock()
	defer tar.Unlock()
	tar.bItems[item.key] = item
}

func (s *segMap) Delete(key string) {
	tar := s.items[s.hashKeyM(key)]
	tar.Lock()
	defer tar.Unlock()
	delete(tar.bItems, key)
}

func (s *segMap) Get(key string) (*item, bool) {
	tar := s.items[s.hashKeyM(key)]
	tar.RLock()
	defer tar.RUnlock()
	if res, has := tar.bItems[key]; has {
		return res, true
	} else {
		return nil, false
	}
}

func NewSegMap(cap int) *segMap {
	bucketSize := 128
	items := make([]*bucket, bucketSize)
	for i := 0; i < bucketSize; i++ {
		items[i] = &bucket{bItems: make(map[string]*item)}
	}
	return &segMap{
		items: items,
	}
}

func (s *segMap) hashKeyM(key string) int {
	if len(key) < 64 {
		var sts [64]byte
		copy(sts[:], key)
		return int(crc32.ChecksumIEEE(sts[:len(key)]) % 128)
	}
	return int(crc32.ChecksumIEEE([]byte(key)) % 128)
}
