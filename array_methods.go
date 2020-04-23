package main

import (
	"math/rand"
	"time"
)

// スライスを分割する
func SliceArr(arr []int, numOfEachSlice int) [][]int {
	arrs := [][]int{}
	len := len(arr)

	for i := 0; i < len; i += numOfEachSlice {
		end := i + numOfEachSlice
		if len < end {
			end = len
		}
		arrs = append(arrs, arr[i:end])
	}

	return arrs
}

// スライスをシャッフルする
func Shuffle(arr []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
