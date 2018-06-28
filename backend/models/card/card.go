package card

type Card struct {
	GeneralWithDeleted `xorm:"extends"`
	UniqueId           string   `xorm:"varchar(32) notnull unique comment('分析唯一id')" form:"uniqueId" json:"uniqueId"`
	Name               string   `xorm:"varchar(128) notnull unique(album_id) comment('卡片名称')" form:"name" json:"name"`
	AlbumId            uint     `xorm:"notnull unique comment('所属卡包id')" form:"albumId" json:"albumId"`
	Size               uint     `xorm:"default(0) comment('卡片大小，单位KB')" form:"size" json:"size"`
	Status             int      `xorm:"smallint(4) default(0) comment('卡片状态:0未编辑1已编辑2已下架3已上架')" form:"status" json:"status"`
	ImgUrl             string   `xorm:"varchar(256) notnull comment('图片地址')" form:"imgUrl" json:"imgUrl"`
	Configs            []Config `xorm:"-"`
	Album              Album    `xorm:"-"`
}

type CardStatus int

const (
	CardStatusNoEdit    CardStatus = 0
	CardStatusEdited    CardStatus = 1
	CardStatusNoRelease CardStatus = 2
	CardStatusReleased  CardStatus = 3
)

type CardServiceModel struct {
	Id          uint                    `json:"id"`
	UniqueId    string                  `json:"uniqueId"`
	Name        string                  `json:"name"`
	AlbumName   string                  `json:"albumName"`
	AlbumSource string                  `json:"albumSource"`
	Resources   [][]MessageServiceModel `json:"resources"`
}

type CardModel struct {
	Card  Card  `xorm:"extends"`
	Album Album `xorm:"extends"`
}
