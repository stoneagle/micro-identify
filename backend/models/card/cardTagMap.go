package card

type CardTagMap struct {
	General `xorm:"extends"`
	CardId  uint `xorm:"not null comment('关联卡片id')"`
	TagId   uint `xorm:"not null comment('关联标签id')"`
	Card    `xorm:"- extends"`
	Tag     `xorm:"- extends"`
}

func (m CardTagMap) TableName() string {
	return "card_tag_map"
}
