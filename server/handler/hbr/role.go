package hbr

type cancer struct {
	border          int
	destructionRate float64
	multiDefence    float64
	multiFragile    float64
}

type nabi struct {
	exskill skill

	strength   int
	agile      int
	fieldBoost float64
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

type skill struct {
	minPower    int
	maxPower    int
	dpBoostRate float64
	hpBoostRate float64
}

// func (n *nabi) GetFinalPower(c cancer) int {}
