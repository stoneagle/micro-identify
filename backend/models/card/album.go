package card

type Album struct {
	GeneralWithDeleted `xorm:"extends"`
	Name               string `xorm:"varchar(128) notnull unique(source) comment('专辑名称')"`
	Source             string `xorm:"varchar(128) notnull unique comment('来源')"`
	Release            bool   `xorm:"default(0) comment('发布状态:0未发布1已发布')"`
	Number             uint   `xorm:"default(0) comment('卡片数量')"`
}

func (a Album) TableName() string {
	return "album"
}
