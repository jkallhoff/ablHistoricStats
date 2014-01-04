package main

import (
//"fmt"
)

func ProcessAverageWOBA(players PlayersList, result *LeagueResult) {
	var woba = make([]float64, 0)
	var (
		greatMark, goodMark, averageMark, badMark, horribleMark                               int64
		numberGreat, numberGood, numberAverage, numberBelowAverage, numberBad, numberHorrible int64
	)

	numberOfPlayers := 0
	for _, p := range players {
		if p.AB > 1 {
			numberOfPlayers = numberOfPlayers + 1
			woba = append(woba, p.WOBA)
		}
	}

	totalWOBA := 0.00
	for _, w := range woba {
		totalWOBA = totalWOBA + w
	}

	averageMark = int64((totalWOBA / float64(numberOfPlayers)) * 1000)
	goodMark = averageMark + 20
	greatMark = goodMark + 20
	badMark = averageMark - 20
	horribleMark = badMark - 20

	for _, p := range players {
		if p.AB > MinimumAtBats {
			intWOBA := int64(p.WOBA * 1000)
			switch {
			case intWOBA >= greatMark:
				numberGreat = numberGreat + 1
				break
			case intWOBA >= goodMark && intWOBA < greatMark:
				numberGood = numberGood + 1
				break
			case intWOBA >= averageMark && intWOBA < goodMark:
				numberAverage = numberAverage + 1
				break
			case intWOBA < averageMark && intWOBA > badMark:
				numberBelowAverage = numberBelowAverage + 1
				break
			case intWOBA <= badMark && intWOBA > horribleMark:
				numberBad = numberBad + 1
				break
			case intWOBA <= horribleMark:
				numberHorrible = numberHorrible + 1
				break
			}
		}
	}

	result.WOBAResult.AverageMark = averageMark
	result.WOBAResult.BadMark = badMark
	result.WOBAResult.GoodMark = goodMark
	result.WOBAResult.GreatMark = greatMark
	result.WOBAResult.HorribleMark = horribleMark
	result.WOBAResult.NumberAverage = numberAverage
	result.WOBAResult.NumberBad = numberBad
	result.WOBAResult.NumberBelowAverage = numberBelowAverage
	result.WOBAResult.NumberGood = numberGood
	result.WOBAResult.NumberGreat = numberGreat
	result.WOBAResult.NumberHorrible = numberHorrible

}
