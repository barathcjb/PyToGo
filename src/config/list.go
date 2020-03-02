package config

import (
	"sync"
)

type List struct {
	items []interface{}
	lock  sync.RWMutex
}

func (list *List) Append(item interface{}) {
	list.lock.RLock()
	defer list.lock.RUnlock()
	if (list.items == nil) || len(list.items) == 0 {
		list.items = []interface{}{}
	}
	list.items = append(list.items, item)
	list.items[len(list.items)-1] = item

}

func (list *List) Remove(item interface{}) interface{} {
	list.lock.Lock()
	defer list.lock.Unlock()
	if (list.items == nil) || len(list.items) == 0 {
		return nil
	}
	for index, element := range list.items {
		if element == item {
			list.items = append(list.items[:index], list.items[index+1:])
		}
	}
	return item
}

func (list *List) Items() []interface{} {
	list.lock.RLock()
	defer list.lock.RUnlock()
	return list.items

}
