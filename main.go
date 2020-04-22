package main

import (
	"fmt"
	"os"
)

func main() {
	allParticipants := 18
	participantsInEachGroup := 6
	repeatCnt := 3
	trials := 20

	if allParticipants%participantsInEachGroup != 0 {
		fmt.Println("このプログラムは参加人数がグループ数で割り切れない場合に対応していません。")
		fmt.Println("割り切れるように数を設定してください。")
		os.Exit(0)
	}

	p := NewParticipants(allParticipants)

	for i := 0; i < trials; i++ {
		sr := NewScoreRecord(allParticipants)

		for i := 0; i < repeatCnt; i++ {
			p.Shuffle()
			sr.Record(p.arr, participantsInEachGroup)
		}

		sd := sr.CalcStandardDeviation()
		if sd < 20 {
			sr.DisplayTable()
			fmt.Println("標準偏差: ", fmt.Sprintf("%.2f", sd))
			fmt.Println()
		}
	}
}
