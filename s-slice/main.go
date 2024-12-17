package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
)

// 回调函数
func ds(a int, f func([]int)) {
	i := []int{a, 1, 2, 3, 8}
	f(i)
}

func main() {
	ds(3, func(s []int) {
		for _, v := range s {
			fmt.Println(v)
		}
	})
	var arr []int16
	for i := 0; i < 4; i++ {
		a := int16(rand.Int() % 100)
		arr = append(arr, a) // 往数组里面添加数据
	}
	for _, val := range arr {
		s := fmt.Sprintf("%4d, %p\n", val, &val)
		fmt.Printf(s)
	}
	// 排序
	slices.Sort(arr)
	for _, val := range arr {
		fmt.Printf("%04d, %v, %p\n", val, val, &val)
	}

	sa, _ := strconv.ParseInt("-3423232232", 10, 64) //将字符串转成int
	fmt.Println(sa)
	sa1, _ := strconv.Atoi("23")
	fmt.Println(sa1)
	fmt.Println(strconv.Itoa(sa1))
	fmt.Println(len(arr), slices.Min(arr), slices.Max(arr))

}
