package main

import (
	"fmt"
	cache "github.com/xukes/go-study/common/cache"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//var map1 map[string]int32
	//map1["sds"] = 2
	map2 := make(map[string]*int32)
	var i int32 = 23
	map2["sds"] = &i

	for k, v := range map2 {
		fmt.Println(k, *v)
	}
	_, isPresent := map2["sds"]

	delete(map2, "s3ds")
	fmt.Println(isPresent)

	cache.Instance().SetVal("name", &Person{Name: "xx", Age: 19})
	cache.Instance().SetVal("user", "user school")
	per := *cache.Instance().GetVal("name").(*Person)
	fmt.Println(per)
	fmt.Println(cache.Instance().GetVal("user"))

}
