package singleFlight

import (
	"fmt"
	"sync"
)

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	m sync.Map
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	if existing, ok := g.m.Load(key); ok {
		c := existing.(*call)
		c.wg.Wait()
		fmt.Println("secondrequest")
		return c.val, c.err
	}
	c := &call{}
	g.m.Store(key, c)
	c.wg.Add(1)
	c.val, c.err = fn()
	c.wg.Done()
	g.m.Delete(key)
	return c.val, c.err
}
