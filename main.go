package main

import (
	"fmt"
	"os"
)

func main() {
	allParticipants := 18
	participantsInEachGroup := 6
	repeatCnt := 3
	trials := 1

	if allParticipants%participantsInEachGroup != 0 {
		fmt.Println("このプログラムは参加人数がグループ数で割り切れない場合に対応していません。")
		fmt.Println("割り切れるように数を設定してください。")
		os.Exit(0)
	}

	pc := NewParticipantCombinations(allParticipants, repeatCnt)

	for i := 0; i < trials; i++ {
		sr := NewScoreRecord(allParticipants)

		for _, combination := range pc.combinations {
			sr.Record(combination, participantsInEachGroup)
		}

		sd := sr.CalcStandardDeviation()
		if sd < 3 {
			fmt.Println("各回の組み合わせ: ")
			for _, combination := range pc.combinations {
				fmt.Printf("%v\n", combination)
			}
			sr.DisplayTable()
			fmt.Println("標準偏差:", fmt.Sprintf("%.2f", sd))
			fmt.Println()
		}
	}
}
