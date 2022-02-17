package core

import (
	"fmt"
	"os"
	"time"

	"github.com/lovelock/gomemcache/v3/memcache"
)

func Get(key, host string, port int) {
	mc := memcache.New(fmt.Sprintf("%s:%d", host, port))
	mc.DisableCAS = true
	mc.Timeout = 1000 * time.Millisecond

	value, err := mc.Get(key)
	if err != nil {
		if err != memcache.ErrCacheMiss {
			fmt.Printf("error occurred: %+v\n", err)
		} else {
			fmt.Println("missed")
		}
		return
	}

	fmt.Println(string(value.Value))
}

func Set(key, value, host string, port int) {
	mc := memcache.New(fmt.Sprintf("%s:%d", host, port))
	mc.Timeout = 1000 * time.Millisecond

	err := mc.Set(&memcache.Item{Key: key, Value: []byte(value)})
	if err != nil {
		fmt.Printf("error ocurred: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("key " + key + " set with value " + value)
}
