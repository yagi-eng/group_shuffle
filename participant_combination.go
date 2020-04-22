package main

import (
	"math/rand"
	"time"
)

// 参加者の配列を管理する型
type ParticipantCombinations struct {
	combinations [][]int
}

// コンストラクタ
func NewParticipantCombinations(allParticipants int, repeatCnt int) *ParticipantCombinations {
	combinations := make([][]int, repeatCnt)

	combination := make([]int, allParticipants)
	for i := range combination {
		combination[i] = i
	}

	for i := 0; i < repeatCnt; i++ {
		shuffle(combination)
		combinations[i] = make([]int, len(combination))
		copy(combinations[i], combination)
	}

	return &ParticipantCombinations{combinations: combinations}
}

// スライスをシャッフルする
func shuffle(arr []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
