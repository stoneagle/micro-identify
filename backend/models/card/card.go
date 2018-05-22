package card

type Card struct {
	GeneralWithDeleted `xorm:"extends"`
	UniqueId           string `xorm:"varchar(32) notnull unique comment('分析唯一id')" form:"uniqueId" json:"uniqueId" binding:"required"`
	Name               string `xorm:"varchar(128) notnull unique(album_id) comment('卡片名称')" form:"name" json:"name" binding:"required"`
	AlbumId            uint   `xorm:"notnull unique comment('所属卡包id')" form:"albumId" json:"albumId" binding:"required"`
	Size               uint   `xorm:"default(0) comment('卡片大小，单位KB')" form:"size" json:"size"`
	Status             int    `xorm:"smallint(4) default(0) comment('卡片状态:0未编辑1已编辑2已下架3已上架')" form:"status" json:"status"`
	ImgUrl             string `xorm:"varchar(256) notnull comment('图片地址')" form:"imgUrl" json:"imgUrl"  binding:"required"`

	Configs []Config `xorm:"-"`
}

func (c Card) TableName() string {
	return "card"
}
