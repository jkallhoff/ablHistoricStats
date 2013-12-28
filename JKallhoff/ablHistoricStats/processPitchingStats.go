package main

func ProcessPitchingStats(players PlayersList, result *LeagueResult) {

	numberOfValidSP := 0
	numberOfValidMR := 0
	numberOfValidCL := 0

	groupedSPFIP := 0.0
	groupedMRFIP := 0.0
	groupedCLFIP := 0.0

	groupedSPBB9 := 0.0
	groupedMRBB9 := 0.0
	groupedCLBB9 := 0.0

	groupedSPK9 := 0.0
	groupedMRK9 := 0.0
	groupedCLK9 := 0.0

	groupedSPBABIP := 0.0
	groupedMRBABIP := 0.0
	groupedCLBABIP := 0.0

	for _, p := range players {
		if p.IP > 0.1 {
			switch {
			case p.POS == "SP":
				groupedSPFIP = groupedSPFIP + p.FIP
				groupedSPBB9 = groupedSPBB9 + p.BB9
				groupedSPK9 = groupedSPK9 + p.K9
				groupedSPBABIP = groupedSPBABIP + p.BABIP
				numberOfValidSP = numberOfValidSP + 1
				break
			case p.POS == "MR":
				groupedMRFIP = groupedMRFIP + p.FIP
				groupedMRBB9 = groupedMRBB9 + p.BB9
				groupedMRK9 = groupedMRK9 + p.K9
				groupedMRBABIP = groupedMRBABIP + p.BABIP
				numberOfValidMR = numberOfValidMR + 1
				break
			case p.POS == "CL":
				groupedCLFIP = groupedCLFIP + p.FIP
				groupedCLBB9 = groupedCLBB9 + p.BB9
				groupedCLK9 = groupedCLK9 + p.K9
				groupedCLBABIP = groupedCLBABIP + p.BABIP
				numberOfValidCL = numberOfValidCL + 1
				break
			}
		}
	}

	result.AverageSPFIP = groupedSPFIP / float64(numberOfValidSP)
	result.AverageSPBB9 = groupedSPBB9 / float64(numberOfValidSP)
	result.AverageSPK9 = groupedSPK9 / float64(numberOfValidSP)
	result.AverageSPBABIP = groupedSPBABIP / float64(numberOfValidSP)

	result.AverageMRFIP = groupedMRFIP / float64(numberOfValidMR)
	result.AverageMRBB9 = groupedMRBB9 / float64(numberOfValidMR)
	result.AverageMRK9 = groupedMRK9 / float64(numberOfValidMR)
	result.AverageMRBABIP = groupedMRBABIP / float64(numberOfValidMR)

	result.AverageCLFIP = groupedCLFIP / float64(numberOfValidCL)
	result.AverageCLBB9 = groupedCLBB9 / float64(numberOfValidCL)
	result.AverageCLK9 = groupedCLK9 / float64(numberOfValidCL)
	result.AverageCLBABIP = groupedCLBABIP / float64(numberOfValidCL)

	result.AverageFIP = (groupedSPFIP + groupedMRFIP + groupedCLFIP) / float64(numberOfValidSP+numberOfValidMR+numberOfValidCL)
	result.AverageBABIP = (groupedSPBABIP + groupedMRBABIP + groupedCLBABIP) / float64(numberOfValidSP+numberOfValidMR+numberOfValidCL)
}
