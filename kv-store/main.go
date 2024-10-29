package main

import (
	"fmt"
	"time"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func main() {
	kv := NewKVStore()

	// Set a kv pair with no TTL
	{
		fmt.Println("Setting key1: value1")
		kv.Set("key1", "value1")

		fmt.Println("(Sleeping)")
		time.Sleep(time.Millisecond * 50)

		fmt.Println("Getting `key1`")
		_, ok := kv.Get("key1")
		fmt.Println("Found:", ok)
		fmt.Println("----------")
	}

	// Get an invalid key
	{
		fmt.Println("Getting `key2`")
		_, ok := kv.Get("key2")
		fmt.Println("Found:", ok)
		fmt.Println("----------")
	}

	// Get an expired key
	{
		fmt.Println("Setting key2: value2 (with TTL 50ms)")
		kv.SetWithTTL("key2", "value2", time.Millisecond*50)

		fmt.Println("(Sleeping)")
		time.Sleep(time.Millisecond * 100)

		fmt.Println("Getting `key2`")
		_, ok := kv.Get("key2")
		fmt.Println("Found:", ok)
		fmt.Println("----------")

	}

	// Get a valid key with TTL
	{
		fmt.Println("Setting key3: value3 (with TTL 50ms)")
		kv.SetWithTTL("key3", "value3", time.Millisecond*50)

		fmt.Println("(Sleeping)")
		time.Sleep(time.Millisecond * 25)

		fmt.Println("Getting `key3`")
		_, ok := kv.Get("key3")
		fmt.Println("Found:", ok)
	}
}
