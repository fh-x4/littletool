package hbr

const typeWeapon = "weapon"
const typeElem = "element"

const (
	cut = iota
	stabs
	blunt

	light
	fire
	dark
	ice
	thunder
)

type cancer struct {
	border          int
	destructionRate float64
	defenceDown     float64
	fragile         float64
	weakPoint       []boostType
	resist          []boostType
}

type boostType struct {
	types     string
	subType   int
	boostRate float64
}

func (c *cancer) CaculateEnemyBoost() float64 {
	return c.destructionRate * c.defenceDown
}

type nabi struct {
	exskill skill

	strength   int
	agile      int
	fieldBoost float64
	// critical
	criticalRate        int // %
	criticalDamageBoost float64
	// boost
	attackBoost float64
	mindBoost   float64
	// hit
	hitNum  int
	hitRate float64
}

type skill struct {
	minPower       int
	maxPower       int
	powerDiff      int
	weightStrength int
	weightAgile    int

	finalPower  int
	dpBoostRate float64
	hpBoostRate float64
	weaponType  int
	weaponElem  int
}

func (n *nabi) CaculateFinalPower(c cancer) {
	weightBorder := (n.exskill.weightStrength*n.strength + n.exskill.weightAgile*n.agile) /
		(n.exskill.weightStrength + n.exskill.weightAgile)
	cancerBorder := c.border
	if n.criticalRate > 100 {
		cancerBorder -= 50
	}

	if weightBorder > cancerBorder {
		if weightBorder-cancerBorder >= n.exskill.powerDiff {
			n.exskill.finalPower = n.exskill.maxPower
		} else {
			n.exskill.finalPower = (n.exskill.maxPower-n.exskill.minPower)*((weightBorder-cancerBorder)/n.exskill.powerDiff) +
				n.exskill.minPower
		}
	} else {
		if cancerBorder-weightBorder >= n.exskill.powerDiff/2 {
			n.exskill.finalPower = 1
		} else {
			n.exskill.finalPower = n.exskill.minPower * (1 - (cancerBorder-weightBorder)/(n.exskill.powerDiff/2))
		}
	}
}

func (n *nabi) CaculateWeakpointBoost(c cancer) float64 {
	var boost float64 = 1
	for _, v := range c.weakPoint {
		if v.types == typeWeapon {
			if v.subType == n.exskill.weaponType {
				boost *= v.boostRate
			}
		}
		if v.types == typeElem {
			if v.subType == n.exskill.weaponElem {
				boost *= v.boostRate
			}
		}
	}
	for _, v := range c.resist {
		if v.types == typeWeapon {
			if v.subType == n.exskill.weaponType {
				boost *= v.boostRate
			}
		}
		if v.types == typeElem {
			if v.subType == n.exskill.weaponElem {
				boost *= v.boostRate
			}
		}
	}
	if boost > 1 {
		boost *= n.mindBoost
		boost *= c.fragile
	}
	return boost
}

func (n *nabi) CaculateFriendlyBoost() float64 {
	base := n.attackBoost * float64(n.hitNum) * n.hitRate * n.criticalDamageBoost * n.fieldBoost
	if n.exskill.hpBoostRate != 1 {
		base *= n.exskill.hpBoostRate
	} else if n.exskill.dpBoostRate != 1 {
		base *= n.exskill.dpBoostRate
	}
	return base
}

// 94363533
// 261396
