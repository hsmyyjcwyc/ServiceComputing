package main

import "fmt"

func partition(array []int, left, right int) int {

	base_number := array[left]
	// 先从右往左找，将第一个小于基数的数放到左边，再从左往右找，将第一个大于基数的数放到右边
	for left < right {
		for (array[right] >= base_number && right > left) {
			right --
		}
		array[left] = array[right]
		for (array[left] <= base_number && left < right) {
			left ++
		}
		array[right] = array[left]
	}
	array[right] = base_number
	return right
}

func quickSort(array []int, left, right int) {
	// 数组为空时直接返回
	if array == nil {
		return
	}
	// 当分区只有一个数时该分区排序完成
	if left >= right {
		return
	}
	index := partition(array,left,right)
	quickSort(array, left, index-1)
	quickSort(array, index+1, right)
}

func main() {
	array1 := []int {}
	array2 := []int {6,33,4,5,1,2}
	fmt.Println(array1)
	quickSort(array1, 0, len(array1)-1)
	fmt.Println(array1)
	fmt.Println(array2)
	quickSort(array2, 0, len(array2)-1)
	fmt.Println(array2)
}