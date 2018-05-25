package card

type Config struct {
	GeneralWithDeleted `xorm:"extends"`
	Name               string    `xorm:"varchar(64) notnull comment('配置集名称') unique(card_id)" form:"name" json:"name"`
	CardId             uint      `xorm:"not null comment('所属卡片id') unique" form:"cardId" json:"cardId"`
	Metadata           string    `xorm:"TEXT comment('扩展配置')" form:"metadata" json:"metadata"`
	Messages           []Message `xorm:"-"`
	Message            Message   `xorm:"extends" json:"-"`
}

func (c Config) TableName() string {
	return "config"
}
