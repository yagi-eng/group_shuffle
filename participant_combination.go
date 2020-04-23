package main

import (
	"fmt"
)

// 参加者の組み合わせを管理する型
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
		Shuffle(combination)
		combinations[i] = make([]int, len(combination))
		copy(combinations[i], combination)
	}

	return &ParticipantCombinations{combinations: combinations}
}

// 参加者の組み合わせを表示する
func (pc *ParticipantCombinations) Display(participantsInEachGroup int) {
	fmt.Println("組み分け:")
	for _, combination := range pc.combinations {
		slicedCombination := SliceArr(combination, participantsInEachGroup)
		fmt.Printf("%v\n", slicedCombination)
	}
}
