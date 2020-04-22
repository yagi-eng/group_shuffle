package main

import (
	"fmt"
	"os"
)

func main() {
	allParticipants := 18
	participantsInEachGroup := 6
	repeatCnt := 3
	trials := 50

	if allParticipants%participantsInEachGroup != 0 {
		fmt.Println("このプログラムは参加人数がグループ数で割り切れない場合に対応していません。")
		fmt.Println("割り切れるように数を設定してください。")
		os.Exit(0)
	}

	p := NewParticipants(allParticipants)

	for i := 0; i < trials; i++ {
		sr := NewScoreRecord(allParticipants)
		participantsTable := make([][]int, repeatCnt)

		for i := 0; i < repeatCnt; i++ {
			p.Shuffle()
			sr.Record(p.arr, participantsInEachGroup)
		}

		sd := sr.CalcStandardDeviation()
		if sd < 1.5 {
			fmt.Println("各回の組み合わせ: ")
			for _, row := range participantsTable {
				fmt.Printf("%v\n", row)
			}
			sr.DisplayTable()
			fmt.Println("標準偏差:", fmt.Sprintf("%.2f", sd))
			fmt.Println()
		}
	}
}
