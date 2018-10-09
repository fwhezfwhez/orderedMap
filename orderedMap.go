package orderedMap

import (
	"sync"
)

type OrderedMap struct{
	M *sync.Mutex
	Map interface{}
	Order []interface{}
}

// new a map
func New(mp func() interface{}) *OrderedMap{
	return &OrderedMap{M: &sync.Mutex{}, Map:mp(),Order: make([]interface{},0,10)}
}

// Set key,value into the map
func (om *OrderedMap)Set(set func()interface{}){
	key := set()
	for _,v :=range om.Order {
		if v == key {
			return
		}
	}
	om.Order = append(om.Order,key)
}

// Get the value by key, returns value,true if value exits,else returns zero,false
func (om *OrderedMap) Get(get func()(interface{},bool)) (interface{},bool){
	return get()
}

// Delete element from the map
func (om *OrderedMap) Delete(del func()interface{}){
	for i,v :=range om.Order{
		if v == del() {
			om.Order = append(om.Order[0:i], om.Order[i+1:]...)
			break
		}
	}
}
