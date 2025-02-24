package hbr

import (
	"context"

	"github.com/fh-x4/littletool/component/httpserver"
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
	} `json:"cancer"`
	Nabi struct {
		Strength            int     `json:"strength"`
		Agile               int     `json:"agile"`
		FieldBoost          float64 `json:"field_boost"`
		CriticalRate        int     `json:"critical_rate"`
		CriticalDamageBoost int     `json:"critical_damage_boost"`
		AttackBoost         float64 `json:"attack_boost"`
		MindBoost           float64 `json:"mind_boost"`
		HitNum              int     `json:"hit_num"`
		HitRate             float64 `json:"hit_rate"`
	} `json:"nabi"`
	Skill struct {
		MinPower    int     `json:"min_power"`
		MaxPower    int     `json:"max_power"`
		DpBoostRate float64 `json:"dp_boost_rate"`
		HpBoostRate float64 `json:"hp_boost_rate"`
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
	return nil
}

type HandlerGen struct{}

func (hg *HandlerGen) GenHandler() httpserver.IHandler {
	return &damageCaculateHandler{
		req: &damageCaculateReq{},
		rsp: &damageCaculateRsp{},
	}
}
