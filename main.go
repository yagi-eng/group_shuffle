package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	// 参加者数
	allParticipants := 36
	// 1グループ毎の人数
	participantsInEachGroup := 6
	// 組み分けの繰り返し回数
	repeatCnt := 3
	// 組み分けの試行回数
	trials := 10000

	if allParticipants%participantsInEachGroup != 0 {
		fmt.Println("このプログラムは参加人数がグループ数で割り切れない場合に対応していません。")
		fmt.Println("割り切れるように数を設定してください。")
		os.Exit(0)
	}

	bettersd := math.MaxFloat64
	var betterPc *ParticipantCombinations
	var betterSr *ScoreRecord
	for i := 0; i < trials; i++ {
		pc := NewParticipantCombinations(allParticipants, repeatCnt)
		sr := NewScoreRecord(allParticipants)

		for _, combination := range pc.combinations {
			sr.Record(combination, participantsInEachGroup)
		}

		tmpSd := sr.CalcStandardDeviation()
		if tmpSd < bettersd {
			bettersd = tmpSd
			betterPc, betterSr = pc, sr
		}
	}

	// 結果を出力
	betterPc.Display(participantsInEachGroup)
	betterSr.Display(repeatCnt)
	fmt.Println("標準偏差:", fmt.Sprintf("%.2f", bettersd))
	fmt.Println()
}
