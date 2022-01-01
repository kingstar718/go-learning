package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {

	c := cache.New(5*time.Minute, 5*time.Minute)

	c.Set("foo", "bar", cache.DefaultExpiration)

	c.Set("baz", 42, cache.DefaultExpiration)

	food, found := c.Get("foo")
	if found {
		fmt.Println(food)
	}
	if baz, found := c.Get("baz"); found {
		fmt.Println(baz)
	}

	if x, found := c.Get("foo"); found {
		foo := x.(string)
		fmt.Println(foo)
	}

	cacheStruct := MyCacheStruct{id: 1, name: "jack", sex: true}
	c.Set("jack", &cacheStruct, cache.DefaultExpiration)
	if x, found := c.Get("jack"); found {
		myCache := x.(*MyCacheStruct)
		fmt.Println(myCache)
	}

}

type MyCacheStruct struct {
	id   int32
	name string
	sex  bool
}
