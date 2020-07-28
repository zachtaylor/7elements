// Code generated by go-gengen(v0.0.1) DO NOT EDIT.

package account

import (
	"sync"
)

// Cache is an observable concurrent in-memory datastore
type Cache struct {
	dat   map[string]*T
	mu    sync.Mutex
	obs   []Func
}

// Storer is an abstraction of in-memory datastore
type Storer interface {
	Get(string) *T
	Set(string, *T)
	Each(Func)
	Sync(func(map[string]*T))
	Keys() []string
	Observe(Func)
	Remove(string)
}

// Func is a callback func
type Func = func(string, *T)

// NewCache returns a new Cache
func NewCache() *Cache {
	return &Cache{
		dat: make(map[string]*T),
		obs:   make([]Func, 0),
	}
}

// Get returns the *T for a string
func (c *Cache) Get(k string) *T { return c.dat[k] }

// Set saves a *T for a string
func (c *Cache) Set(k string, v *T) {
	c.mu.Lock()
	if v != nil {
		c.dat[k] = v
	} else {
		delete(c.dat, k)
	}
	c.mu.Unlock()
	for _, f := range c.obs {
		f(k, v)
	}
}

// Each calls the func for each string,*T in this Cache
func (c *Cache) Each(f Func) {
	c.mu.Lock()
	for k, v := range c.dat {
		f(k, v)
	}
	c.mu.Unlock()
}

// Sync calls the func within the cache lock state
func (c *Cache) Sync(f func(map[string]*T)) {
	c.mu.Lock()
	f(c.dat)
	c.mu.Unlock()
}

// Keys returns a new slice with all the string keys
func (c *Cache) Keys() []string {
	c.mu.Lock()
	keys := make([]string, 0, len(c.dat))
	for k := range c.dat {
		keys = append(keys, k)
	}
	c.mu.Unlock()
	return keys
}

// Observe adds a func to be called when a *T is explicitly set
func (c *Cache) Observe(f Func) { c.obs = append(c.obs, f) }

// Remove deletes a string,*T
func (c *Cache) Remove(k string) { c.Set(k, nil) }
