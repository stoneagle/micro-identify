package card

type Release struct {
	General `xorm:"extends"`
	AlbumId uint   `xorm:"notnull comment('所属卡包id')" form:"albumId" json:"albumId" binding:"required"`
	AgentId string `xorm:"notnull varchar(64) comment('发布渠道id')" form:"agentId" json:"agentId" binding:"required"`
}

func (r Release) TableName() string {
	return "release"
}
