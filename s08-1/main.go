package main

import "fmt"

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

	f := func() {
		fmt.Println("defer")
	}
	defer f()
}
