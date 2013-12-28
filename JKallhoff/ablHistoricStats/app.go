package main

import (
	"flag"
	"fmt"
)

var Players = make(PlayersList, 0)
var (
	MinimumAtBats int64
)

type Processor func(PlayersList, *LeagueResult)
type ProcessorList []Processor

func main() {
	var battingFile string
	var pitchingFile string
	MinimumAtBats = 10

	flag.StringVar(&battingFile, "b", "", "The relative path to the batting information")
	flag.StringVar(&pitchingFile, "p", "", "The relative path to the pitching information")

	flag.Parse()

	loadBattingSamples(battingFile)

	processors := ProcessorList{ProcessAverageWOBA, ProcessPitchingStats, ProcessBattingWOBAMaps}
	result := processors.Run(Players)

	fmt.Println("wOBA Results")
	fmt.Println("===========================")
	fmt.Printf("   Great wOBA: %10d\n", result.WOBAResult.GreatMark)
	fmt.Printf("    Good wOBA: %10d\n", result.WOBAResult.GoodMark)
	fmt.Printf(" Average wOBA: %10d\n", result.WOBAResult.AverageMark)
	fmt.Printf("     Bad wOBA: %10d\n", result.WOBAResult.BadMark)
	fmt.Printf("Horrible wOBA: %10d\n", result.WOBAResult.HorribleMark)

	fmt.Println("\n\n")

	fmt.Println("FIP Results")
	fmt.Println("===========================")
	fmt.Printf("   Average FIP: %10.2f\n", result.AverageFIP)
	fmt.Printf("Average SP FIP: %10.2f\n", result.AverageSPFIP)
	fmt.Printf("Average MR FIP: %10.2f\n", result.AverageMRFIP)
	fmt.Printf("Average CL FIP: %10.2f\n", result.AverageCLFIP)

	fmt.Println("\n\n")

	fmt.Println("BABIP Results")
	fmt.Println("===========================")
	fmt.Printf("   Average BABIP: %10.3f\n", result.AverageBABIP)
	fmt.Printf("Average SP BABIP: %10.3f\n", result.AverageSPBABIP)
	fmt.Printf("Average MR BABIP: %10.3f\n", result.AverageMRBABIP)
	fmt.Printf("Average CL BABIP: %10.3f\n", result.AverageCLBABIP)

	fmt.Println("\n\n")

	fmt.Println("BB/9 Results")
	fmt.Println("===========================")
	fmt.Printf("Average SP BB/9: %10.2f\n", result.AverageSPBB9)
	fmt.Printf("Average MR BB/9: %10.2f\n", result.AverageMRBB9)
	fmt.Printf("Average CL BB/9: %10.2f\n", result.AverageCLBB9)

	fmt.Println("\n\n")

	fmt.Println("K/9 Results")
	fmt.Println("===========================")
	fmt.Printf("Average SP K/9: %10.2f\n", result.AverageSPK9)
	fmt.Printf("Average MR K/9: %10.2f\n", result.AverageMRK9)
	fmt.Printf("Average CL K/9: %10.2f\n", result.AverageCLK9)

	fmt.Printf("wOBA: %#v", *(result.WOBAResult))
}

func (l ProcessorList) Run(players PlayersList) *LeagueResult {
	result := NewLeagueResult()
	for _, p := range l {
		p(players, result)
	}
	return result
}
