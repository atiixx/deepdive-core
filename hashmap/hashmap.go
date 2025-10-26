package hashmap

import (
	"fmt"
	"hash/fnv"
)

type OpenHashMap[V any] struct {
	buckets       []*MapNode[V]
	numOfElements int
}

type ClosedHashMap[V any] struct {
	buckets       [][]*MapNode[V]
	numOfElements int
}

func NewClosedHashMap[V any]() ClosedHashMap[V] {
	return ClosedHashMap[V]{buckets: make([][]*MapNode[V], 10)}
}

func NewOpenHashMap[V any]() OpenHashMap[V] {
	return OpenHashMap[V]{buckets: make([]*MapNode[V], 10)}
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
	m.numOfElements += 1
	fmt.Println("Added Element {", key, ",", value, "} to Closed Map.")
	return true
}

func (m *ClosedHashMap[V]) removeClosed(key string) bool {
	index := int(hash(key)) % len(m.buckets)
	indexSlice := m.buckets[index]
	for i, element := range indexSlice {
		if element.Key == key {
			indexSlice = append(indexSlice[:i], indexSlice[i+1:]...)
			m.numOfElements -= 1
			fmt.Println("Removed Element {", key, ",", element.Value, "} from Closed Map.")
			return true
		}
	}
	return false
}

func (m ClosedHashMap[V]) getClosed(key string) V {
	index := int(hash(key)) % len(m.buckets)
	indexSlice := m.buckets[index]
	for _, element := range indexSlice {
		if element.Key == key {
			return element.Value
		}
	}
	var zero V
	return zero
}

func (m *ClosedHashMap[V]) resizeClosed() {
	newBuckets := make([][]*MapNode[V], len(m.buckets)*2)
	newMap := &ClosedHashMap[V]{buckets: newBuckets}
	for _, element := range m.buckets {
		for _, kvpair := range element {
			newMap.insertClosed(kvpair.Key, kvpair.Value)
		}
	}
	fmt.Println("Closed HashMap resized from ", len(m.buckets), " to ", len(m.buckets)*2, " elements.")
	m = newMap
}

func (m *OpenHashMap[V]) insertOpen(key string, value V) bool {
	index := int(hash(key)) % len(m.buckets)
	for ; m.buckets[index:] != nil; index += 1 {
		if index == len(m.buckets) {
			index = 0
		}
	}

	m.buckets[index] = &MapNode[V]{Key: key, Value: value}
	m.numOfElements += 1
	if float32(m.numOfElements/len(m.buckets)) > 0.5 {
		m.resizeOpen()
	}
	fmt.Println("Added Element {", key, ",", value, "} to Closed Map.")
	return true
}

func (m *OpenHashMap[V]) removeOpen(key string) bool {
	index := int(hash(key)) % len(m.buckets)
	for ; m.buckets[index:] != nil; index += 1 {
		if m.buckets[index].Key == key {
			m.buckets[index].active = false
			return true
		}
	}
	return false
}

func (m OpenHashMap[V]) getOpen(key string) V {
	index := int(hash(key)) % len(m.buckets)
	for ; m.buckets[index:] != nil; index += 1 {
		if m.buckets[index].Key == key && m.buckets[index].active {
			return m.buckets[index].Value
		}
	}
	var zero V
	return zero
}

func (m *OpenHashMap[V]) resizeOpen() {
	newBuckets := make([]*MapNode[V], len(m.buckets)*2)
	newMap := &OpenHashMap[V]{buckets: newBuckets}
	for _, element := range m.buckets {
		newMap.insertOpen(element.Key, element.Value)
	}
	fmt.Println("Open HashMap resized from ", len(m.buckets), " to ", len(m.buckets)*2, " elements.")
	m = newMap
}
