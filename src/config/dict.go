package config

import "sync"

type Dictionary struct {
	item map[interface{}]interface{}
	lock sync.RWMutex
}

func (dict *Dictionary) Set(key, value interface{}) {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	if dict.item == nil {
		dict.item = make(map[interface{}]interface{})
	}
	dict.item[key] = value
}

func (dict *Dictionary) Get(key interface{}) interface{} {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	return dict.item[key]
}

func (dict *Dictionary) Keys() []interface{} {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	keys := []interface{}{}
	for i, _ := range dict.item {
		keys = append(keys, i)
	}
	return keys
}

func (dict *Dictionary) Values() []interface{} {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	values := []interface{}{}
	for i, _ := range dict.item {
		values = append(values, dict.item[i])
	}
	return values
}

func (dict *Dictionary) Delete(key interface{}) bool {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	for i, _ := range dict.item {
		if key == i {
			delete(dict.item, key)
			return true
		}
	}
	return false
}

func (dict *Dictionary) Clear() {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	dict.item = make(map[interface{}]interface{})
}

func (dict *Dictionary) Items() []interface{} {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	items := []interface{}{}
	for key, value := range dict.item {
		items = append(items, []interface{}{key, value})
	}
	return items
}
