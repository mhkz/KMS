package main

import "fmt"


func main ()  {
	source := []int{1, 2, 10, 9, 7, 5, 4,8,6,3}
	bubbleSort(&source)
}

/**
* 冒泡排序
*/

func bubbleSort (s*[]int) {
	for i :=0; i < len(*s) - 1; i ++ {
		fmt.Printf("ssss", s*[i])
	}
}