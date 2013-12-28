package main

import (
	"code.google.com/p/go.net/html"
	"log"
	"os"
	"strconv"
	"strings"
)

type playerSamplerFunc func(*html.Node)

var leagueBeingProcessed string

func loadBattingSamples(league string) {
	htmlPath := league
	leagueBeingProcessed = league

	file, err := os.Open(htmlPath)
	if err != nil {
		log.Fatalln("Could not open the batting file")
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatalln("Could not close the batting file")
		}
	}(file)

	node, _ := html.Parse(file)
	processNode(pullTogetherBatterSample, node)

	return
}

func processNode(playerSampler playerSamplerFunc, node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "tr" {
		playerSampler(node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		processNode(playerSampler, child)
	}
}

func pullTogetherBatterSample(node *html.Node) {
	player := new(Player)
	iterator := 1
	for column := node.FirstChild; column != nil; column = column.NextSibling {
		if column.FirstChild != nil {
			switch iterator {
			case 2:
				player.POS = strings.TrimSpace(column.FirstChild.Data)
			case 4:
				player.Name = strings.TrimSpace(column.FirstChild.Data)
			case 6:
				player.Age = fetchInteger(column.FirstChild)
			case 8:
				player.Bats = column.FirstChild.Data
			case 10:
				player.Throws = column.FirstChild.Data
			case 12:
				overall := column.FirstChild.Data
				player.Overall, _ = strconv.ParseInt(overall[:1], 10, 8)
			case 14:
				potential := column.FirstChild.Data
				player.Potential, _ = strconv.ParseInt(potential[:1], 10, 8)
			case 16:
				player.WE = column.FirstChild.Data
			case 18:
				player.INT = column.FirstChild.Data
			case 20:
				player.Contact = fetchInteger(column.FirstChild)
			case 22:
				player.Gap = fetchInteger(column.FirstChild)
			case 24:
				player.Power = fetchInteger(column.FirstChild)
			case 26:
				player.Eye = fetchInteger(column.FirstChild)
			case 28:
				player.AvoidsStrikes = fetchInteger(column.FirstChild)
			case 30:
				player.Stuff = fetchInteger(column.FirstChild)
			case 32:
				player.Movement = fetchInteger(column.FirstChild)
			case 34:
				player.Control = fetchInteger(column.FirstChild)
			case 36:
				player.Stamina = fetchInteger(column.FirstChild)
			case 38:
				player.Speed = fetchInteger(column.FirstChild)
			case 40:
				player.Stealing = fetchInteger(column.FirstChild)
			case 42:
				player.BaseRunning = fetchInteger(column.FirstChild)
			case 44:
				player.AB = fetchInteger(column.FirstChild)
			case 46:
				player.HR = fetchInteger(column.FirstChild)
			case 48:
				player.AVG = fetchFloat(column.FirstChild)
			case 50:
				player.OBP = fetchFloat(column.FirstChild)
			case 52:
				player.SLG = fetchFloat(column.FirstChild)
			case 54:
				player.WOBA = fetchFloat(column.FirstChild)
			case 56:
				player.OPSPlus = fetchInteger(column.FirstChild)
			case 58:
				player.BVORP = fetchFloat(column.FirstChild)
			case 60:
				player.BWAR = fetchFloat(column.FirstChild)
			case 62:
				player.IP = fetchFloat(column.FirstChild)
			case 64:
				player.ERA = fetchFloat(column.FirstChild)
			case 66:
				player.BABIP = fetchFloat(column.FirstChild)
			case 68:
				player.WHIP = fetchFloat(column.FirstChild)
			case 70:
				player.HR9 = fetchFloat(column.FirstChild)
			case 72:
				player.BB9 = fetchFloat(column.FirstChild)
			case 74:
				player.K9 = fetchFloat(column.FirstChild)
			case 76:
				player.FIP = fetchFloat(column.FirstChild)
			case 78:
				player.PVORP = fetchFloat(column.FirstChild)
			case 80:
				player.PWAR = fetchFloat(column.FirstChild)
			case 82:
				player.PPG = fetchFloat(column.FirstChild)
			case 84:
				player.DEFF = fetchFloat(column.FirstChild)
			}
		}
		iterator++
	}

	if player.Name != "" && player.Name != "Name" {
		Players = append(Players, player)
	}
}

func fetchInteger(node *html.Node) (num int64) {
	num, _ = strconv.ParseInt(node.Data, 10, 16)
	return
}

func fetchFloat(node *html.Node) (num float64) {
	num, _ = strconv.ParseFloat(node.Data, 64)
	return
}
