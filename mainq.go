package main

import (
	"fmt"
	"sync"
)

type cache interface {
	set()
	get()
	delete()
}

type cacheImpl struct {
	data map[string]interface{}
}

func main() {
	c := &cacheImpl{data: make(map[string]interface{})}
	//c.set("key", "value")
	//fmt.Println(c.get("key"))
	//c.delete("key")
	//fmt.Println(c.get("key"))

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		c.set("q", "0")
		c.set("a", "1")
		c.set("b", "2")
		c.delete("a")
		c.set("c", "3")

		wg.Done()
	}()

	go func() {
		fmt.Println(c.get("q"))
		fmt.Println(c.get("a"))
		fmt.Println(c.get("b"))
		fmt.Println(c.get("c"))

		wg.Done()
	}()

	wg.Wait()
}

func (c *cacheImpl) set(key string, value interface{}) {
	c.data[key] = value
}

func (c *cacheImpl) get(key string) interface{} {
	return c.data[key]
}

func (c *cacheImpl) delete(key string) {
	delete(c.data, key)
}
