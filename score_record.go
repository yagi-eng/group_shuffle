package main

import (
	"fmt"
	"math"
)

// 重複度を管理する型
type ScoreRecord struct {
	scores     []int
	countTable [][]int
}

// コンストラクタ
func NewScoreRecord(len int) *ScoreRecord {
	scores := make([]int, len)
	countTable := createTableFilledZero(len)
	return &ScoreRecord{scores: scores, countTable: countTable}
}

// 全ての要素が0のテーブルを生成する
func createTableFilledZero(len int) [][]int {
	table := make([][]int, len)
	for i := 0; i < len; i++ {
		table[i] = make([]int, len)
	}
	return table
}

// スコアと同席回数を記録する
func (sr *ScoreRecord) Record(participants []int, participantsInEachGroup int) {
	groups := sliceArr(participants, participantsInEachGroup)

	for _, group := range groups {
		sr.recordEachGroup(group)
	}
}

func sliceArr(participants []int, participantsInEachGroup int) [][]int {
	groups := [][]int{}
	len := len(participants)

	for i := 0; i < len; i += participantsInEachGroup {
		end := i + participantsInEachGroup
		if len < end {
			end = len
		}
		groups = append(groups, participants[i:end])
	}

	return groups
}

func (sr *ScoreRecord) recordEachGroup(group []int) {
	len := len(group)
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if i == j {
				continue
			}

			x, y := group[i], group[j]
			// 同席回数を記録
			sr.countTable[x][y]++
			// スコアを記録
			sr.scores[x] += sr.countTable[x][y]
		}
	}
}

// 標準偏差を計算する
func (sr *ScoreRecord) CalcStandardDeviation() float64 {
	sum := 0
	for _, v := range sr.scores {
		sum += v
	}
	len := len(sr.scores)
	ave := float64(sum) / float64(len)

	numerator := 0.0
	for _, v := range sr.scores {
		numerator += math.Pow(float64(v)-ave, 2)
	}
	return math.Sqrt(numerator / float64(len))
}

// テーブルの中身と要素の内訳を表示する
func (sr *ScoreRecord) Display() {
	fmt.Println("同席回数をカウントしたテーブル: ")

	cnt := make([]int, 4)
	for _, row := range sr.countTable {
		fmt.Printf("%v\n", row)

		for _, v := range row {
			cnt[v]++
		}
	}

	for i, v := range cnt {
		fmt.Printf("%dの数: %d\n", i, v)
	}
}
