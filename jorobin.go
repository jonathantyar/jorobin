package jorobin

import (
	"errors"
	"sync/atomic"
)

var ErrServersNotExists = errors.New("servers dose not exist")

type RoundRobin interface {
	Next() string       // Return next item
	Total() int         // Return count of items
	New(newItem string) // Add new item to roundrobin
	Remove(item string) // Remove an item from the roundrobin
	Clear()             // Clear all items from the roundrobin
}

type roundrobin struct {
	items []string
	next  uint32
}

func New(items ...string) (RoundRobin, error) {
	if len(items) == 0 {
		return nil, ErrServersNotExists
	}

	return &roundrobin{
		items: items,
	}, nil
}

func NewWithArray(items []string) (RoundRobin, error) {
	if len(items) == 0 {
		return nil, ErrServersNotExists
	}

	return &roundrobin{
		items: items,
	}, nil
}

func (r *roundrobin) Next() string {
	n := atomic.AddUint32(&r.next, 1)
	return r.items[(int(n)-1)%len(r.items)]
}

func (r *roundrobin) Total() int {
	return len(r.items)
}

func (r *roundrobin) New(newItem string) {
	r.items = append(r.items, newItem)
}

func (r *roundrobin) Remove(item string) {
	for i, v := range r.items {
		if v == item {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return
		}
	}
}

func (r *roundrobin) Clear() {
	r.items = []string{}
}
