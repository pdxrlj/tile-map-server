package tile_rules

type TileRules struct {
	Path string
}

type TileRulesInterface interface {
	Rules() string
}

func NewTileRules() *TileRules {
	rule := &TileRules{}
	return rule
}

// Rules 返回瓦片存储地址
func (t *TileRules) Rules(rules TileRulesInterface) string {
	return rules.Rules()
}
