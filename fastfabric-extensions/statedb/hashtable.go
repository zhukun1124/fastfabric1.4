package statedb

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
)

// ValueHashtable the set of Items
type ValueHashtable struct {
	items map[string][]byte
	lock  sync.RWMutex
}

func NewHT() *ValueHashtable {
	return &ValueHashtable{lock:sync.RWMutex{}}
}

func (ht *ValueHashtable) Put(key []byte, value []byte) error {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	if ht.items == nil {
		ht.items = make(map[string][]byte)
	}
	ht.items[string(key)] = value
	fmt.Println("Write key:",string(key))
	return nil
}

// Remove item with key k from hashtable
func  (ht *ValueHashtable) Remove(key []byte) error {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	delete(ht.items, string(key))
	return nil
}

// Get item with key k from the hashtable
func (ht *ValueHashtable) Get(key []byte) ([]byte, error) {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	if val, ok := ht.items[string(key)]; ok {
		return val, nil
	} else {
		return nil, errors.New("key not found")
	}
}

// Size returns the number of the hashtable elements
func (ht *ValueHashtable) Size() int {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return len(ht.items)
}

func (ht *ValueHashtable) Cleanup() {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	ht.items = nil
}

func (ht *ValueHashtable) GetKeys(sk []byte, ek []byte)[][]byte {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	keys := make([][]byte, len(ht.items))
	if len(sk)!= 0 && sk[len(sk)-1] == 0x01{
		sk = sk[:len(sk)-1]
	}

	if len(ek)!= 0 && ek[len(ek)-1] == 0x01{
		ek = ek[:len(ek)-1]
		ek[len(ek)-1] += 1
	}
	fmt.Println("hastable start:", string(sk), ", end:", string(ek))

	i := 0
	for k := range ht.items {
		fmt.Println("key int table:", k)
		x := []byte(k)
		var toCompare []byte
		if (len(x)> len(ek)){
			toCompare = x[:len(ek)]
		}else{
			toCompare = x
		}
		if bytes.Compare(sk, toCompare) < 1 && bytes.Compare(toCompare, ek) < 1 {
			keys[i] = x
			i++
		}
	}
	return keys[:i]
}
