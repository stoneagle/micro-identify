package card

type Message struct {
	General  `xorm:"extends"`
	ConfigId uint   `xorm:"not null comment('所属配置集id')" form:"configId" json:"configId" binding:"required"`
	Type     int    `xorm:"not null default(0) comment('类别:0文本1音频')" form:"type" json:"type" binding:"required"`
	Order    int    `xorm:"not null default(0) comment('排列顺序:0顺位最前')" form:"order" json:"order" binding:"required"`
	Detail   string `xorm:"TEXT comment('描述内容')" form:"detail" json:"detail" binding:"required"`
}

func (m Message) TableName() string {
	return "message"
}
