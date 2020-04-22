package main

import (
	"math/rand"
	"time"
)

// 参加者の配列を管理する型
type Participants struct {
	arr []int
}

// コンストラクタ
func NewParticipants(allParticipants int) *Participants {
	arr := make([]int, allParticipants)
	for i := range arr {
		arr[i] = i
	}
	return &Participants{arr: arr}
}

// 参加者の配列をシャッフルする
func (p *Participants) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range p.arr {
		j := rand.Intn(i + 1)
		p.arr[i], p.arr[j] = p.arr[j], p.arr[i]
	}
}
