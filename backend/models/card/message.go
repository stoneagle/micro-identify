package card

type Message struct {
	General  `xorm:"extends"`
	ConfigId uint   `xorm:"not null comment('所属配置集id')" form:"configId" json:"configId"`
	Type     int    `xorm:"not null default(0) comment('类别:0文本1音频')" form:"type" json:"type"`
	Order    int    `xorm:"not null default(0) comment('排列顺序:0顺位最前')" form:"order" json:"order"`
	Detail   string `xorm:"TEXT comment('描述内容')" form:"detail" json:"detail"`
}

func (m Message) TableName() string {
	return "message"
}
