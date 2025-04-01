package main

import (
	"fmt"
	"github.com/xukes/go-study/common/cache"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 将函数作为val存储在map中。个人觉得这有点太鸡肋了。想不出这个有啥实际作用。
	map1 := make(map[string]func(string, string) int, 20)
	map1["sds"] = func(s string, s2 string) int {
		return 2
	}
	_, ok := map1["sds"]
	fmt.Println(ok)
	if _, ok := map1["sds"]; ok {
		// 如果存在这个key,则进入该流程
	}
	delete(map1, "sds") // 删除map1中的sds

	map2 := make(map[string]*int32)
	var i int32 = 23
	map2["sds"] = &i

	for k, v := range map2 {
		fmt.Println(k, *v)
	}
	_, isPresent := map2["sds"]

	delete(map2, "s3ds")
	fmt.Println(isPresent)

	cache.Ins.SetVal("name", &Person{Name: "xx", Age: 19})
	cache.Ins.SetVal("user", "user school")
	per := *cache.Ins.GetVal("name").(*Person)
	fmt.Println(per)
	fmt.Println(cache.Ins.GetVal("user"))

}

type Info struct {
	mu sync.Mutex
}

func Update(info *Info) {
	info.mu.Lock()
	defer info.mu.Unlock()

}
