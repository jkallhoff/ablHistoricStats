package main

type PlayersList []*Player

type Player struct {
	POS, Name, LVL, Bats, Throws, WE, INT                                                      string
	Age, Overall, Potential, Contact, Gap, Power, Eye, AvoidsStrikes, Stuff, Movement, Control int64
	Stamina, Speed, Stealing, BaseRunning, AB, HR, OPSPlus                                     int64
	GroundFly, AVG, OBP, SLG, WOBA, BVORP, BWAR, BABIP, DEFF                                   float64
	ERA, WHIP, HR9, BB9, K9, FIP, PVORP, PWAR, PPG, ZR, IP                                     float64
}
