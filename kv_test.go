// This file has tests just for the simple GET/PUT.

package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

// Try to Get a non-exist key
// Check if the return kv is nil
func TestGetWrongKey(t *testing.T) {
	key := "123"
	value := "hello"
	Put(key, value)

	t1 := Get("456")
	if t1 != nil {
		t.Errorf("Expect Get to return nil when accessing an non-exisit key: %+v\n", t1)
	}
}

func testGetThroughputConfigSingle(ch chan int64, key string, nrGet int) {
	start := time.Now()
	for i := 0; i < nrGet; i++ {
		kv := Get(key)
		if kv == nil {
			ch <- -1
		}
	}
	latency_ns := time.Since(start).Nanoseconds()
	latency_ns /= int64(nrGet)
	ch <- latency_ns
}

func testGetThroughputConfig(key string, nrGet int, nrGoroutines int) (float64, error) {
	if key == "" || nrGet <= 0 {
		return 0, errors.New("Invalid Parameter")
	}

	ch := make(chan int64)
	start := time.Now()
	for i := 0; i < nrGoroutines; i++ {
		go testGetThroughputConfigSingle(ch, key, nrGet)
	}
	for i := 0; i < nrGoroutines; i++ {
		ret := <-ch
		if ret < 0 {
			return 0, errors.New("testThroughputSingle failed")
		}
		fmt.Printf("  Latency: %d ns\n", ret)

	}
	latency_s := time.Since(start).Seconds()
	tput := float64(nrGet*nrGoroutines) / latency_s
	return tput, nil
}

// Test Get's Throughput on various configurations
func TestGetThroughput(t *testing.T) {
	key := "testKey"
	value := "testValue"
	Put(key, value)

	nrGet := 1000000
	nrGoroutines := []int{1, 2, 4, 8, 16}
	for _, nrGo := range nrGoroutines {
		t1, err := testGetThroughputConfig(key, nrGet, nrGo)
		if err != nil {
			t.Errorf("Fail to test Get Throughput\n")
		}
		fmt.Printf("nrGoroutines=%d Throughput %f IOPS\n", nrGo, t1)
	}
}

func TestDumpHashTable(t *testing.T) {
	key := "123"
	value := "hello"
	Put(key, value)
	DumpHashTable(os.Stdout)
}
