package hashmap

import (
	"hash/fnv"
)

type OpenHashMap[V any] struct {
	buckets       [10]*MapNode[V]
	numOfElements int
}

type ClosedHashMap[V any] struct {
	buckets       [10][]*MapNode[V]
	numOfElements int
}

func hash(key string) uint64 {
	hash := fnv.New64()
	hash.Write([]byte(key))
	return hash.Sum64()
}

func (m *ClosedHashMap[V]) insertClosed(key string, value V) bool {
	index := int(hash(key)) % len(m.buckets)
	indexSlice := m.buckets[index]
	for _, element := range indexSlice {
		if element.Key == key {
			return false
		}
	}
	m.buckets[index] = append(m.buckets[index], &MapNode[V]{Key: key, Value: value})
	if float32(m.numOfElements/len(m.buckets)) > 0.5 {
		m.resizeClosed()
	}
	return true
}

func (m *ClosedHashMap[V]) removeClosed(key string) {}

func (m ClosedHashMap[V]) getClosed(key string) V {}

func (m *ClosedHashMap[V]) resizeClosed() {}

func (m *OpenHashMap[V]) insertOpen(key string, value V) bool {
	index := int(hash(key)) % len(m.buckets)

	if float32(m.numOfElements/len(m.buckets)) > 0.5 {
		m.resizeOpen()
	}
	return true
}

func (m *OpenHashMap[V]) removeOpen(key string) {}

func (m OpenHashMap[V]) getOpen(key string) V {}

func (m *OpenHashMap[V]) resizeOpen() {}
