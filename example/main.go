package main

import (
	"fmt"
	omp "orderedMap"
	// omp "github.com/fwhezfwhez/orderdMap"
)

func main() {
	// init, specific the map type as map[string]int
	om := omp.New(func() interface{} {
		m := make(map[string]int)
		return m
	})

	// set key,value
	om.Set(func() interface{} {
		om.Map.(map[string]int)["ft"] = 1
		return "ft"
	})

	// get key,value
	v,ok:=om.Get(func()(interface{},bool){
		v,ok :=om.Map.(map[string]int)["ft"]
		return v,ok
	})
	fmt.Println("whether has this key:", ok)
	fmt.Println("its value is:", v)

	// print its order
	fmt.Println("before delete:", om.Order)
	// delete by key
	om.Delete(func()interface{}{
		delete(om.Map.(map[string]int),"ft")
		return "ft"
	})

	// print its order
	fmt.Println("after delete:", om.Order)
}
