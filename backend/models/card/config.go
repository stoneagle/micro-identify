package card

type Config struct {
	GeneralWithDeleted `xorm:"extends"`
	Name               string `xorm:"varchar(64) notnull comment('配置集名称') unique(card_id)"`
	CardId             uint   `xorm:"not null comment('所属卡片id') unique"`
	Metadata           string `xorm:"TEXT comment('扩展配置')"`
}

func (c Config) TableName() string {
	return "config"
}
