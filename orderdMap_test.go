package orderedMap

import (
	"fmt"
	"reflect"
	"testing"
)
var om *OrderedMap
func Init(){
	om = New(func() interface{}{
		m :=make(map[string]int)
		return m
	})
	om.Set(func()interface{}{
		om.Map.(map[string]int)["ft"] = 5
		return "ft"
	})
}

func TestOrderedMap_Set(t *testing.T) {
	Init()
	om.Set(func()interface{}{
		om.Map.(map[string]int)["ft"] = 5
		return "ft"
	})
	fmt.Println(om)
}

func TestOrderedMap_Get(t *testing.T) {
	Init()
	e,ok:=om.Get(func()(interface{},bool){
		e,ok := om.Map.(map[string]int)["ft"]
		return e,ok
	})
	fmt.Println(ok)
	fmt.Println(e)
	fmt.Println(reflect.TypeOf(e).String())

	e2,ok2:=om.Get(func()(interface{},bool){
		e,ok := om.Map.(map[string]int)["ft2"]
		return e,ok
	})
	fmt.Println(ok2)
	fmt.Println(e2)
	fmt.Println(reflect.TypeOf(e2).String())
}

func TestOrderedMap_Delete(t *testing.T) {
	Init()
	om.Set(func()interface{}{
		om.Map.(map[string]int)["ft4"] = 7
		return "ft4"
	})
	fmt.Println("before delete", om.Order)

	om.Delete(func()interface{}{
		delete(om.Map.(map[string]int),"ft4")
		return "ft4"
	})
	fmt.Println("after delete", om.Order)
}

func TestOrderedMap_DoubleSet(t *testing.T){
	Init()
	om.Set(func()interface{}{
		om.Map.(map[string]int)["ft4"] = 7
		return "ft4"
	})

	fmt.Println("before double set:", om.Order)
	om.Set(func()interface{}{
		om.Map.(map[string]int)["ft4"] = 8
		return "ft4"
	})

	om.Set(func()interface{}{
		om.Map.(map[string]int)["ft"] = 8
		return "ft4"
	})
	fmt.Println("after double set:", om.Order)
}

