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

			// スコアを記録
			// 自分自身の数字がスコアには反映されないため、
			// 自分より大きい数字の人を記録する際は-1する
			if x < y {
				sr.scores[x] += y - 1
			} else {
				sr.scores[x] += y
			}

			// 同席回数を記録
			sr.countTable[x][y]++
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

	tmpSum := 0.0
	for _, v := range sr.scores {
		tmpSum += math.Pow(float64(v)-ave, 2)
	}

	return math.Sqrt(tmpSum / float64(len))
}

// （デバッグ用）テーブルの中身を表示する
func (sr *ScoreRecord) DisplayTable() {
	countZero := 0
	countThree := 0
	for _, row := range sr.countTable {
		fmt.Printf("%v\n", row)

		for _, v := range row {
			if v == 0 {
				countZero++
			} else if v == 3 {
				countThree++
			}
		}
	}

	fmt.Println("0の数: ", countZero)
	fmt.Println("3の数: ", countThree)
}
