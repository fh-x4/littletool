package hbr

import (
	"context"

	"github.com/fh-x4/littletool/component/httpserver"
	"github.com/fh-x4/littletool/component/logger"
)

type damageCaculateHandler struct {
	req *damageCaculateReq
	rsp *damageCaculateRsp
}
type damageCaculateReq struct {
	Cancer struct {
		Border          int     `json:"border"`
		DestructionRate float64 `json:"destruction_rate"`
		MultiDefence    float64 `json:"multi_defence"`
		MultiFragile    float64 `json:"multi_fragile"`
		WeakPoint       []struct {
			Types     string  `json:"types"`
			SubTypes  int     `json:"sub_type"`
			BoostRate float64 `json:"boost_rate"`
		} `json:"weak_point"`
		Resist []struct {
			Types     string  `json:"types"`
			SubTypes  int     `json:"sub_type"`
			BoostRate float64 `json:"boost_rate"`
		} `json:"resist"`
	} `json:"cancer"`
	Nabi struct {
		Strength            int     `json:"strength"`
		Agile               int     `json:"agile"`
		FieldBoost          float64 `json:"field_boost" default:"1"`
		CriticalRate        int     `json:"critical_rate" default:"100"`
		CriticalDamageBoost float64 `json:"critical_damage_boost" default:"1.5"`
		AttackBoost         float64 `json:"attack_boost" default:"1"`
		MindBoost           float64 `json:"mind_boost" default:"1"`
		HitNum              int     `json:"hit_num" default:"0"`
		HitRate             float64 `json:"hit_rate" default:"0"`
	} `json:"nabi"`
	Skill struct {
		MinPower       int     `json:"min_power"`
		MaxPower       int     `json:"max_power"`
		PowerDiff      int     `json:"power_diff"`
		WeightStrength int     `json:"weight_strength"`
		WeightAgile    int     `json:"weight_agile"`
		DpBoostRate    float64 `json:"dp_boost_rate" default:"1"`
		HpBoostRate    float64 `json:"hp_boost_rate" default:"1"`
		WeaponType     int     `json:"weapon_type"`
		WeaponElem     int     `json:"weapon_elem"`
	} `json:"skill"`
}
type damageCaculateRsp struct {
	Damage int `json:"damage"`
}

func (h *damageCaculateHandler) GetRequest() interface{} {
	return h.req
}
func (h *damageCaculateHandler) GetRespond() interface{} {
	return h.rsp
}
func (h *damageCaculateHandler) Call(ctx context.Context) httpserver.IError {
	enemy := cancer{
		border:          h.req.Cancer.Border,
		destructionRate: h.req.Cancer.DestructionRate,
		defenceDown:     h.req.Cancer.MultiDefence,
		fragile:         h.req.Cancer.MultiFragile,
	}
	friendly := nabi{
		exskill: skill{
			minPower:       h.req.Skill.MinPower,
			maxPower:       h.req.Skill.MaxPower,
			powerDiff:      h.req.Skill.PowerDiff,
			weightStrength: h.req.Skill.WeightStrength,
			weightAgile:    h.req.Skill.WeightAgile,
			dpBoostRate:    h.req.Skill.DpBoostRate,
			hpBoostRate:    h.req.Skill.HpBoostRate,
			weaponType:     h.req.Skill.WeaponType,
			weaponElem:     h.req.Skill.WeaponElem,
		},
		strength:            h.req.Nabi.Strength,
		agile:               h.req.Nabi.Agile,
		fieldBoost:          h.req.Nabi.FieldBoost,
		criticalRate:        h.req.Nabi.CriticalRate,
		criticalDamageBoost: h.req.Nabi.CriticalDamageBoost,
		attackBoost:         h.req.Nabi.AttackBoost,
		mindBoost:           h.req.Nabi.MindBoost,
		hitNum:              h.req.Nabi.HitNum,
		hitRate:             h.req.Nabi.HitRate,
	}
	for _, v := range h.req.Cancer.WeakPoint {
		enemy.weakPoint = append(enemy.weakPoint, boostType{
			types:     v.Types,
			subType:   v.SubTypes,
			boostRate: v.BoostRate,
		})
	}
	for _, v := range h.req.Cancer.Resist {
		enemy.resist = append(enemy.resist, boostType{
			types:     v.Types,
			subType:   v.SubTypes,
			boostRate: v.BoostRate,
		})
	}
	if !enemy.CheckValid() || !friendly.CheckValid() {
		logger.GetLogger().Infof("input invalid:%v,%v", enemy.CheckValid(), friendly.CheckValid())
		return nil
	}

	friendly.CaculateFinalPower(enemy)
	weakpoint := friendly.CaculateWeakpointBoost(enemy)
	buff := friendly.CaculateFriendlyBoost()
	debuff := enemy.CaculateEnemyBoost()
	h.rsp.Damage = int(friendly.exskill.finalPower * weakpoint * buff * debuff)
	logger.GetLogger().Infof("finalPower=%v,weakpoint=%v,buff=%v,debuff=%v", friendly.exskill.finalPower, weakpoint, buff, debuff)

	return nil
}

type HandlerGen struct{}

func (hg *HandlerGen) GenHandler() httpserver.IHandler {
	return &damageCaculateHandler{
		req: &damageCaculateReq{},
		rsp: &damageCaculateRsp{},
	}
}
