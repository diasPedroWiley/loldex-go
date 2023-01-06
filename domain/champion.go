package domain

type Stats struct {
	HP                   float64 `json:"hp"`
	HPPerLevel           float64 `json:"hpperlevel"`
	MP                   float64 `json:"mp"`
	MPPerLevel           float64 `json:"mpperlevel"`
	MoveSpeed            float64 `json:"movespeed"`
	Armor                float64 `json:"armor"`
	ArmorPerLevel        float64 `json:"armorperlevel"`
	SpellBlock           float64 `json:"spellblock"`
	SpellBlockPerLevel   float64 `json:"spellblockperlevel"`
	AttackRange          float64 `json:"attackrange"`
	HPRegen              float64 `json:"hpregen"`
	HPRegenPerLevel      float64 `json:"hpregenperlevel"`
	MPRegen              float64 `json:"mpregen"`
	MPRegenPerLevel      float64 `json:"mpregenperlevel"`
	Crit                 float64 `json:"crit"`
	CritPerLevel         float64 `json:"critperlevel"`
	AttackDamage         float64 `json:"attackdamage"`
	AttackDamagePerLevel float64 `json:"attackdamageperlevel"`
	AttackSpeedPerLevel  float64 `json:"attackspeedperlevel"`
	AttackSpeed          float64 `json:"attackspeed"`
}

type Image struct {
	Full   string
	Sprite string
	Group  string
	X      int64
	Y      int64
	W      int64
	H      int64
}

type Spell struct {
	Id          string
	Name        string
	Description string
	Tooltip     string
	LevelTip    struct {
		Label  []string
		Effect []string
	}
	MaxRank      int64
	Cooldown     []float64
	CooldownBurn string
	Cost         []int64
	CostBurn     string
	Effect       [][]int64
	EffectBurn   []string
	CostType     string
}

type BaseChampion struct {
	Id      string `json:"id"`
	Key     string `json:"key"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Image   `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Blurb   string   `json:"blurb"`
	Stats   `json:"stats"`
	Info    struct {
		Attack     int64 `json:"attack"`
		Defense    int64 `json:"defense"`
		Magic      int64 `json:"magic"`
		Difficulty int64 `json:"difficulty"`
	} `json:"info"`
}

type Champion struct {
	BaseChampion
	Skins []struct {
		Id      string
		Num     int64
		Name    string
		Chromas bool
	}
	Lore      string
	AllTips   []string
	EnimyTips []string
	Passive   struct {
		Name        string
		Description string
		Image
	}
}
