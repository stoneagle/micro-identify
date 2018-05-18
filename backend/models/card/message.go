package card

type Message struct {
	General  `xorm:"extends"`
	ConfigId uint   `xorm:"not null comment('所属配置集id')"`
	Type     uint   `xorm:"not null default(0) comment('类别:0文本1音频')"`
	Order    uint   `xorm:"not null default(0) comment('排列顺序:0顺位最前')"`
	Detail   string `xorm:"TEXT comment('描述内容')"`
}

func (m Message) TableName() string {
	return "message"
}
