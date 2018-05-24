package card

type Album struct {
	GeneralWithDeleted `xorm:"extends"`
	Name               string `xorm:"varchar(128) notnull unique(source) comment('专辑名称')" form:"name" json:"name"`
	Source             string `xorm:"varchar(128) notnull unique comment('来源')" form:"source" json:"source"`
	Release            int    `xorm:"smallint(1) default(0) comment('发布状态:0未发布1已发布')" form:"release" json:"release"`
	Number             uint   `xorm:"default(0) comment('卡片数量')" form:"number" json:"number"`
}

func (a Album) TableName() string {
	return "album"
}
