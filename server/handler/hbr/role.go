package hbr

type cancer struct {
	border          int
	destructionRate float64
	multiDefence    float64
	multiFragile    float64
}

type nabi struct {
	strength   int
	agile      int
	minPower   int
	maxPower   int
	fieldBoost float64
	// 大招词条
	dpBoostRate float64
	hpBoostRate float64
	// critical
	criticalRate        int // %
	criticalDamageBoost int
	// boost
	attackBoost float64
	mindBoost   float64
	// hit
	hitNum  int
	hitRate float64
}

// func (n *nabi) GetFinalPower(c cancer) int {}
