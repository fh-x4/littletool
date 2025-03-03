package hbr

const typeWeapon = "weapon"
const typeElem = "element"

const (
	unknown = iota
	cut
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

func (c *cancer) CheckValid() bool {
	if c.border == 0 || c.destructionRate == 0 || c.defenceDown == 0 || c.fragile == 0 {
		return false
	}
	return true
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

func (n *nabi) CheckValid() bool {
	if !n.exskill.CheckValid() {
		return false
	}
	if n.strength == 0 || n.agile == 0 || n.criticalRate < 0 {
		return false
	}
	return true
}

func (n *nabi) CaculateFinalPower(c cancer) {
	weightBorder := (n.exskill.weightStrength*n.strength + n.exskill.weightAgile*n.agile) /
		(n.exskill.weightStrength + n.exskill.weightAgile)
	cancerBorder := c.border
	if n.criticalRate >= 100 {
		cancerBorder -= 50
	}

	if weightBorder > cancerBorder {
		if weightBorder-cancerBorder >= n.exskill.powerDiff {
			n.exskill.finalPower = float64(n.exskill.maxPower)
		} else {
			n.exskill.finalPower = float64(n.exskill.maxPower-n.exskill.minPower)*(float64(weightBorder-cancerBorder)/float64(n.exskill.powerDiff)) +
				float64(n.exskill.minPower)
		}
	} else {
		if cancerBorder-weightBorder >= n.exskill.powerDiff/2 {
			n.exskill.finalPower = 1
		} else {
			n.exskill.finalPower = float64(n.exskill.minPower) * (1 - float64(cancerBorder-weightBorder)/float64(n.exskill.powerDiff/2))
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
	base := n.attackBoost * n.fieldBoost
	// 以现在的版本来说应该都是必暴击的
	if n.criticalRate >= 100 {
		base *= n.criticalDamageBoost
	} else {
		base *= (n.criticalDamageBoost-1)*float64(n.criticalRate/100) + 1
	}
	if n.hitNum > 0 {
		base *= 1 + (float64(n.hitNum) * n.hitRate)
	}
	if n.exskill.hpBoostRate != 1 {
		base *= n.exskill.hpBoostRate
	} else if n.exskill.dpBoostRate != 1 {
		base *= n.exskill.dpBoostRate
	}
	return base
}

type skill struct {
	minPower       int
	maxPower       int
	powerDiff      int
	weightStrength int
	weightAgile    int

	finalPower  float64
	dpBoostRate float64
	hpBoostRate float64
	weaponType  int
	weaponElem  int
}

func (s *skill) CheckValid() bool {
	if s.minPower == 0 || s.maxPower == 0 || s.powerDiff == 0 || s.weightStrength == 0 || s.weightAgile == 0 {
		return false
	}
	if s.weaponType == unknown {
		return false
	}
	return true
}
