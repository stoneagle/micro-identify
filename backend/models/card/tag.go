package card

type Tag struct {
	GeneralWithDeleted `xorm:"extends"`
	Name               string `xorm:"varchar(128) notnull unique comment('标签名称')" form:"name" json:"name"`
}

func (t Tag) TableName() string {
	return "tag"
}
