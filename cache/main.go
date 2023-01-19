package main

import "time"

type Cache struct {
	store map[string]*Value
	head  *Value
	tail  *Value
}

type Value struct {
	val      int
	lastUsed time.Time
	next     *Value
	prev     *Value
}

func (c *Cache) get(key string) int {
	val, ok := c.store[key]
	if !ok {
		return -1
	}
	val.lastUsed = time.Now()
	val.prev.next = val.next

	val.next = c.head
	c.head.prev = val
	val.prev = nil

	c.head = val

	return val.val
}

func (c *Cache) put(key string, val int) {
	_, ok := c.store[key]
	if !ok {

	}
	v := &Value{val: val, lastUsed: time.Now(), next: c.head}
	c.head.prev = v
	c.head = val
}

func main() {

}
