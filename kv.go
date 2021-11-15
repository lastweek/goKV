package main

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type KeyValueEntry struct {
	hashKey uint32
	key     string
	value   string
	next    *KeyValueEntry

	nrHit int
}

type KeyValueHead struct {
	next      *KeyValueEntry
	lock      sync.Mutex
	nrChained int
}

const hashTableSize = 16

var nrKVEntries = 0

var hashTable [hashTableSize]KeyValueHead

func hashStringToInt(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func Put(key string, value string) *KeyValueEntry {
	// Create the new entry
	hash := hashStringToInt(key)
	kv := KeyValueEntry{}
	kv.key = key
	kv.value = value
	kv.hashKey = hash

	index := hash % hashTableSize
	head := &hashTable[index]

	// Add to the hash bucket
	head.lock.Lock()
	if head.next == nil {
		head.next = &kv
		head.nrChained = 1
	} else {
		tmp := head.next
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = &kv
		head.nrChained += 1
	}
	nrKVEntries++
	head.lock.Unlock()

	return &kv
}

func DumpHashTable() {
	fmt.Printf("Total number of KV Entries: %d\n", nrKVEntries)
	for i := 0; i < hashTableSize; i++ {
		head := &hashTable[i]
		if head.nrChained != 0 {
			fmt.Printf("bucket %d    nrChained %d\n", i, head.nrChained)
			kv := head.next
			for kv != nil {
				// fmt.Printf("    keyHash: %x key: %s value: %s\n",
				// 	kv.hashKey, kv.key, kv.value)
				fmt.Printf("%+v\n", kv)
				kv = kv.next
			}
		}
	}
}

func Get(key string) *KeyValueEntry {
	hash := hashStringToInt(key)
	index := hash % hashTableSize
	head := &hashTable[index]

	head.lock.Lock()
	defer head.lock.Unlock()

	kv := head.next
	for kv != nil {
		if kv.key == key {
			kv.nrHit++
			return kv
		}
		kv = kv.next
	}
	return nil
}

func main() {
	fmt.Println("Welcome to the GoKV")
}
