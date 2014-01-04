package main

type wobaResult struct {
	GreatMark, GoodMark, AverageMark, BadMark, HorribleMark                               int64
	NumberGreat, NumberGood, NumberAverage, NumberBelowAverage, NumberBad, NumberHorrible int64
}

type LeagueResult struct {
	WOBAResult                                                            *wobaResult
	BattingWOBAMaps                                                       map[int64][]float64
	BattingWARMaps                                                        warItems
	AverageFIP, AverageBABIP                                              float64
	AverageSPFIP, AverageSPBB9, AverageSPK9, AverageSPBABIP, AverageSPWAR float64
	AverageMRFIP, AverageMRBB9, AverageMRK9, AverageMRBABIP, AverageMRWAR float64
	AverageCLFIP, AverageCLBB9, AverageCLK9, AverageCLBABIP               float64
}

func NewLeagueResult() *LeagueResult {
	result := new(LeagueResult)
	result.WOBAResult = new(wobaResult)
	result.BattingWOBAMaps = make(map[int64][]float64)
	return result
}
