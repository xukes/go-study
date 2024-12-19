package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
给一个24小时制的时间（00:00-23:59)，其中有一个或多个数字是问号。问号处可以用任何一个数字代替，问可以表示的最大时间是多少
*/
func Test_1871(t *testing.T) {
	fmt.Println(MaximumMoment("2?:??"))
}
func MaximumMoment(time string) string {
	var arr [4]int
	idx := 0
	for _, v := range strings.Split(time, "") {
		if v != ":" {
			if v == "?" {
				arr[idx] = -1
			} else {
				val, _ := strconv.Atoi(v)
				arr[idx] = val
			}
			idx++
		}
	}
	if arr[0] == -1 && arr[1] < 4 {
		arr[0] = 2
	} else if arr[0] == -1 {
		arr[0] = 1
	}
	if arr[1] == -1 && arr[0] == 1 || arr[1] == -1 && arr[0] == 0 {
		arr[1] = 9
	} else if arr[1] == -1 && arr[0] == 2 {
		arr[1] = 3
	}

	if arr[2] == -1 {
		arr[2] = 5
	}
	if arr[3] == -1 {
		arr[3] = 9
	}
	val := fmt.Sprintf("%v%v:%v%v", arr[0], arr[1], arr[2], arr[3])
	return val
}

// 从一个范围内选择最多整数
/**
给定一个正整数数组 banned 和两个整数 size 和 maxSum，按照以下要求选择一些整数：

可以选择范围在 [1, size] 之内的整数
每个整数 最多 选择一次
不允许选择值在 banned 中的整数
选择的所有整数之和不大于 maxSum
*/
func Test_3856(t *testing.T) {
	fmt.Println(MaxCount([]int{2, 3, 6}, 6, 5))
}
func MaxCount(banned []int, size int, maxSum int) int {

	for i := 0; i < size; i++ {

	}
	return 0
}
